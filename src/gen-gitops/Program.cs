using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Serialization;

namespace gen_gitops
{
    class Program
    {
        private static readonly List<string> Regions = new();
        private static readonly List<District> Districts = new();
        private static readonly List<Store> Stores = new();

        private static readonly JsonSerializerOptions SerializerOptions = new()
        {
            PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
            WriteIndented = true,
            DefaultIgnoreCondition = JsonIgnoreCondition.WhenWritingNull,
        };

        const string environment = "demo";
        const string deploy = "deploy";
        const string domainName = "cseretail.com";

        private static bool GenerateSsl = false;

        public static int Main(string[] args)
        {
            if (Directory.Exists(deploy))
            {
                return LogError("Please delete deploy directory and run again");
            }

            if (args == null || args.Length == 0 || args.Contains("-h") || args.Contains("--h"))
            {
                return Usage();
            }

            if (args.Contains("--ssl"))
            {
                GenerateSsl = true;
            }

            if (ReadFile(args[0]))
            {
                GenerateExpandTargets();
                GenerateTargets();
                GenerateScripts();
            }

            return 0;
        }

        private static bool ReadFile(string fileName)
        {
            try
            {
                if (!File.Exists(fileName))
                {
                    LogError($"File not found: {fileName}");
                    return false;
                }

                string txt = File.ReadAllText(fileName);

                if (string.IsNullOrWhiteSpace(txt))
                {
                    LogError("File is empty or whitespace");
                    return false;
                }

                string[] lines = txt.Split(new char[] { '\n' }, StringSplitOptions.RemoveEmptyEntries | StringSplitOptions.TrimEntries);

                if (lines.Length == 0)
                {
                    LogError("File is empty or whitespace");
                    return false;
                }

                string[] cols;
                HashSet<string> regions = new();
                HashSet<string> districts = new();
                HashSet<string> stores = new();

                foreach (string line in lines)
                {
                    // skip comments
                    if (line[0] !='#')
                    {
                        cols = line.Split(new char[] { '\t' });

                        if (cols.Length >= 4)
                        {
                            Store st = new()
                            {
                                Region = cols[0],
                                State = cols[1],
                                City = cols[2],
                                Number = int.Parse(cols[3]),
                            };

                            if (cols.Length > 4)
                            {
                                st.AzureRegion = cols[4];
                            }

                            if (!stores.Contains(st.Name))
                            {
                                stores.Add(st.Name);
                                Stores.Add(st);

                                if (!districts.Contains(st.District))
                                {
                                    districts.Add(st.District);
                                    Districts.Add(new()
                                    {
                                        Region = st.Region,
                                        State = st.State,
                                        City = st.City,
                                    });
                                }

                                if (!regions.Contains(st.Region))
                                {
                                    regions.Add(st.Region);
                                    Regions.Add(st.Region);
                                }
                            }
                        }
                    }
                }

                // sort the lists
                Regions.Sort();
                Districts.Sort((x, y) => { return string.Compare(x.Name, y.Name); });
                Stores.Sort((x, y) => { return string.Compare(x.Name, y.Name); });
            }
            catch (Exception ex)
            {
                LogError(ex.Message);
                return false;
            }

            return true;
        }

        private static int Usage()
        {
            Console.WriteLine($"Usage: gen-gitops TSVFileName");
            Console.WriteLine();

            return 1;
        }

        private static void GenerateExpandTargets()
        {
            if (!Directory.Exists(deploy))
            {
                Directory.CreateDirectory(deploy);
            }

            Directory.SetCurrentDirectory(deploy);

            Dictionary<string, List<string>> expando = new();
            expando.Add("all", Regions);

            const string file = "expandTargets.json";

            foreach (string r in Regions)
            {
                List<string> districts = new ();
                expando.Add(r, districts);

                foreach (District d in Districts)
                {
                    if (d.Region == r)
                    {
                        districts.Add(d.Name);
                        List<string> stores = new ();
                        expando.Add(d.Name, stores);

                        foreach (Store s in Stores)
                        {
                            if (s.Region == r && s.District == d.Name)
                            {
                                stores.Add(s.Name);
                            }
                        }
                    }
                }
            }

            File.WriteAllText(file, JsonSerializer.Serialize(expando, SerializerOptions));
            Directory.SetCurrentDirectory("..");
        }

        private static void GenerateScripts()
        {
            const string header = "#!/bin/bash\n\n";
            const string cd = "# change to this directory\ncd $(dirname \"${BASH_SOURCE[0]}\")\ncd ..\n\n";
            const string ssl = $" -z {domainName} --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem";

            Directory.CreateDirectory("scripts");
            Directory.SetCurrentDirectory("scripts");

            string create = $"{header}{cd}";
            string delete = $"{header}";
            string curl = $"{header}";

            foreach (Store s in Stores)
            {
                create += $"create-cluster {s.Region} {s.State} {s.City} {s.Number} -l {s.AzureRegion}";

                if (GenerateSsl)
                {
                    create += ssl;
                }

                create += " &\n";

                delete += $"akdc delete {s.Name}\n";
                curl += $"curl https://{s.Name}.{domainName}/tinybench/17; echo \"  {s.Name}\" &\n";
            }

            delete += "\n\n# remove IPs\nrm $(dirname \"${ BASH_SOURCE[0]}\"/ips)\n";

            // save the files
            File.WriteAllText("create.sh", create);
            File.WriteAllText("delete.sh", delete);

            // this will only work if ssl was generated
            if (GenerateSsl)
            {
                File.WriteAllText("curl-all.sh", curl);
            }

            Directory.SetCurrentDirectory("..");
        }

        private static void GenerateTargets()
        {
            if (!Directory.Exists(deploy))
            {
                Directory.CreateDirectory(deploy);
            }

            Directory.SetCurrentDirectory(deploy);

            Config cfg;
            string path;
            string file;

            foreach (string r in Regions)
            {
                path = Path.Combine(".", $"{r}");
                file = Path.Combine(path, "config.json");

                cfg = new Config() { Region = r, Zone = r, Environment = environment };

                Directory.CreateDirectory(path);
                File.WriteAllText(file, JsonSerializer.Serialize(cfg, SerializerOptions));
            }

            foreach (District d in Districts)
            {
                path = Path.Combine(".", $"{d.Name}");
                file = Path.Combine(path, "config.json");

                cfg = new Config() { Region = d.Region, District = d.Name, Zone = d.Name, Environment = environment };

                Directory.CreateDirectory(path);
                File.WriteAllText(file, JsonSerializer.Serialize(cfg, SerializerOptions));
            }

            string flux = GetFluxCheck();

            foreach (Store s in Stores)
            {
                path = Path.Combine(".", $"{s.Name}");
                file = Path.Combine(path, "config.json");

                cfg = new Config() { Region = s.Region, District = s.District, Store = s.Name, Domain = $"{s.Name}.{domainName}", Zone = s.District, Environment = environment };

                Directory.CreateDirectory(path);
                File.WriteAllText(file, JsonSerializer.Serialize(cfg, SerializerOptions));

                // create the flux-check namespace
                path = Path.Combine(path, "flux-check");
                Directory.CreateDirectory(path);
                File.WriteAllText(Path.Combine(path, "namespace.yaml"), flux);
            }

            Directory.SetCurrentDirectory("..");
        }

        private static int LogError(string msg)
        {
            Console.ForegroundColor = ConsoleColor.Red;
            Console.Error.WriteLine(msg);
            Console.ResetColor();
            return 1;
        }

        private static string GetFluxCheck()
        {
            return @"apiVersion: v1
kind: Namespace
metadata:
  labels:
    name: flux-check
  name: flux-check
";
        }
    }

    public class Config
    {
        public string Environment { get; set; }
        public string Region { get; set; }
        public string Zone { get; set; }
        public string District { get; set; }
        public string Store { get; set; }
        public string Domain { get; set; }
    }

    public class District
    {
        public string Region { get; set; }
        public string State { get; set; }
        public string City { get; set; }
        public string Name => $"{Region}-{State}-{City}";

        public override string ToString() => Name;
    }

    public class Store
    {
        public string Region { get; set; }
        public string State { get; set; }
        public string City { get; set; }
        public int Number { get; set; }
        public string AzureRegion { get; set; } = "centralus";
        public string District => $"{Region}-{State}-{City}";
        public string Name => $"{District}-{Number}";

        public override string ToString() => Name;
    }
}
