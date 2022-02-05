using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Serialization;

namespace gen_gitops
{
    class Program
    {
        public static List<string> Regions;
        public static List<District> Districts;
        public static List<Store> Stores;
        public static JsonSerializerOptions SerializerOptions;

        const string deploy = "deploy";
        const string domainName = "cseretail.com";

        public static int Main()
        {
            if (Directory.Exists(deploy))
            {
                return LogError("Please delete deploy directory and run again");
            }

            SerializerOptions = new()
            {
                PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
                WriteIndented = true,
                DefaultIgnoreCondition = JsonIgnoreCondition.WhenWritingNull,
            };
            
            LoadData();
            GenerateTargets();
            GenerateScripts();

            return 0;
        }

        private static void GenerateScripts()
        {
            const string shebang = "#!/bin/bash\n\n";
            const string cd = "cd ..\n\n";

            Directory.CreateDirectory("scripts");
            Directory.SetCurrentDirectory("scripts");
            Directory.CreateDirectory("delete");

            string file;
            string txt;

            foreach (string r in Regions)
            {
                txt = $"{shebang}{cd}";

                foreach (Store s in Stores)
                {
                    if (s.Region == r)
                    {
                        txt += $"./create-cluster.sh {s.Region} {s.State} {s.City} {s.Number} &\n";
                    }
                }

                file = $"{r}.sh";
                File.WriteAllText(file, txt);
            }

            foreach (District d in Districts)
            {
                txt = $"{shebang}{cd}";

                foreach (Store s in Stores)
                {
                    if (s.Region == d.Region && s.State == d.State)
                    {
                        txt += $"./create-cluster.sh {s.Region} {s.State} {s.City} {s.Number} &\n";
                    }
                }

                file = $"{d.State}.sh";
                File.WriteAllText(file, txt);
            }

            foreach (string r in Regions)
            {
                txt = $"{shebang}";

                foreach (Store s in Stores)
                {
                    if (s.Region == r)
                    {
                        txt += $"az group delete -y --no-wait -g {s.Region}-{s.State}-{s.City}-{s.Number}\n";
                    }
                }

                file = Path.Combine("delete", $"{r}.sh");
                File.WriteAllText(file, txt);
            }

            foreach (District d in Districts)
            {
                txt = $"{shebang}";

                foreach (Store s in Stores)
                {
                    if (s.Region == d.Region && s.State == d.State)
                    {
                        txt += $"az group delete -y --no-wait -g {s.Region}-{s.State}-{s.City}-{s.Number}\n";
                    }
                }

                file = Path.Combine("delete", $"{d.State}.sh");
                File.WriteAllText(file, txt);
            }

            Directory.SetCurrentDirectory("..");
        }

        private static void GenerateTargets()
        {
            Directory.CreateDirectory(deploy);
            Directory.SetCurrentDirectory(deploy);

            Config cfg;
            string path;
            string file;

            foreach (string r in Regions)
            {
                path = Path.Combine(".", $"{r}");
                file = Path.Combine(path, "config.json");

                cfg = new Config() { Region = r, Zone = r };

                Directory.CreateDirectory(path);
                File.WriteAllText(file, JsonSerializer.Serialize(cfg, SerializerOptions));

                // create the flux-check namespace
                File.WriteAllText(Path.Combine(path, "check-flux.yaml"), GetFluxCheck());
            }

            foreach (District d in Districts)
            {
                path = Path.Combine(".", $"{d.Name}");
                file = Path.Combine(path, "config.json");

                cfg = new Config() { Region = d.Region, District = d.Name, Zone = d.Name };

                Directory.CreateDirectory(path);
                File.WriteAllText(file, JsonSerializer.Serialize(cfg, SerializerOptions));
            }

            foreach (Store s in Stores)
            {
                path = Path.Combine(".", $"{s.Name}");
                file = Path.Combine(path, "config.json");

                cfg = new Config() { Region = s.Region, District = s.District, Store = s.Name, Domain = $"{s.Name}.{domainName}", Zone = s.District };

                Directory.CreateDirectory(path);
                File.WriteAllText(file, JsonSerializer.Serialize(cfg, SerializerOptions));
            }

            Directory.SetCurrentDirectory("..");
        }

        private static void LoadData()
        {
            Regions = new() { "central", "east", "west" };

            Districts = new()
            {
                new() { Region = "central", State = "tx", City = "austin" },
                new() { Region = "central", State = "tx", City = "dallas" },
                new() { Region = "central", State = "tx", City = "houston" },
                new() { Region = "central", State = "tx", City = "san" },
                new() { Region = "central", State = "tx", City = "north" },
                new() { Region = "central", State = "tx", City = "south" },
                new() { Region = "central", State = "tx", City = "east" },
                new() { Region = "central", State = "tx", City = "west" },

                new() { Region = "central", State = "mo", City = "stlouis" },
                new() { Region = "central", State = "mo", City = "kc" },
                new() { Region = "central", State = "mo", City = "columbia" },
                new() { Region = "central", State = "mo", City = "north" },
                new() { Region = "central", State = "mo", City = "south" },
                new() { Region = "central", State = "mo", City = "east" },
                new() { Region = "central", State = "mo", City = "west" },

                new() { Region = "east", State = "ga", City = "atlanta" },
                new() { Region = "east", State = "ga", City = "athens" },
                new() { Region = "east", State = "ga", City = "north" },
                new() { Region = "east", State = "ga", City = "south" },

                new() { Region = "east", State = "nc", City = "charlotte" },
                new() { Region = "east", State = "nc", City = "raleigh" },
                new() { Region = "east", State = "nc", City = "durham" },
                new() { Region = "east", State = "nc", City = "east" },
                new() { Region = "east", State = "nc", City = "west" },

                new() { Region = "west", State = "ca", City = "la" },
                new() { Region = "west", State = "ca", City = "sfo" },
                new() { Region = "west", State = "ca", City = "sd" },
                new() { Region = "west", State = "ca", City = "sac" },
                new() { Region = "west", State = "ca", City = "north" },
                new() { Region = "west", State = "ca", City = "south" },
                new() { Region = "west", State = "ca", City = "east" },
                new() { Region = "west", State = "ca", City = "west" },

                new() { Region = "west", State = "wa", City = "seattle" },
                new() { Region = "west", State = "wa", City = "spokane" },
                new() { Region = "west", State = "wa", City = "olympia" },
                new() { Region = "west", State = "wa", City = "west" },
                new() { Region = "west", State = "wa", City = "central" },
                new() { Region = "west", State = "wa", City = "east" },
            };

            Stores = new();

            foreach (District d in Districts)
            {
                for (int i = 101; i <= 105; i++)
                {
                    Stores.Add(new Store() { Region = d.Region, State = d.State, City = d.City, Number = i });
                }
            }
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
        public string Region { get; set; }
        public string District { get; set; }
        public string Store { get; set; }
        public string Zone { get;set; }
        public string Environment { get; set; } = "dev";
        public string Domain { get; set; }
    }

    public class District
    {
        public string Region { get; set; }
        public string State { get; set; }
        public string City { get; set; }
        public string Name => $"{Region}-{State}-{City}";
    }

    public class Store
    {
        public string Region { get; set; }
        public string State { get; set; }
        public string City { get; set; }
        public int Number { get; set; }
        public string District => $"{Region}-{State}-{City}";
        public string Name => $"{District}-{Number}";
    }
}
