// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// NewAppCmd creates a new application
var NewAppCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new application",
}

// NewDotnetAppCmd creates a new dotnet application
var NewDotNetAppCmd = &cobra.Command{
	Use:   "dotnet",
	Short: "Create a new dotnet 6 application",
}

func init() {
	NewAppCmd.AddCommand(NewDotNetAppCmd)
	boa.AddScriptCommand(NewDotNetAppCmd, "webapi", "Create a new dotnet WebAPI", fltDotnetWebapiScript())
	boa.AddScriptCommand(NewDotNetAppCmd, "console", "Create a new dotnet console application", fltDotnetConsoleScript())
}

func fltDotnetWebapiScript() string {

	return `
#!/bin/bash

# update the app name if a valid name
export APP_NAME=$1
export APP_LOWER=$(echo "$APP_NAME" | tr '[:upper:]' '[:lower:]')

if [[ "$APP_NAME" =~ ^[A-Z][A-Za-z][A-Za-z][A-Za-z][A-Za-z]+$ ]]
then
    if [ -d "$APP_LOWER" ]
    then
        echo "Directory $APP_LOWER already exists"
        exit 1
    fi

    git pull
    git clone https://github.com/retaildevcrews/dotnet-webapi-template "$APP_LOWER"
    cd "$APP_LOWER" || exit

    rm -rf .devcontainer
    rm -rf .git
    rm -rf .github
    rm -f .gitignore
    rm -f LICENSE
    rm -f curl.sh

    mv src/csapp.csproj "src/$APP_LOWER.csproj"

    sed -i "s/csapp/$APP_LOWER/g" .kic/commands/app/build
    sed -i "s/csapp/$APP_LOWER/g" .kic/commands/app/deploy
    sed -i "s/cd \"\$REPO_BASE\" || exit//g" .kic/commands/app/build
    sed -i "s~deploy/apps~apps~g" autogitops/autogitops.json
    sed -i "s/csapp/$APP_LOWER/g" Dockerfile
    find . -name '*.*' -type f -exec sed -i "s/CSApp/$APP_NAME/g" {} \;
    find . -name '*.*' -type f -exec sed -i "s/csapp/$APP_LOWER/g" {} \;
    dotnet restore src

    git pull
    git add .
    git commit -m "added testapp"
else
    echo "Invalid App Name $1"
fi

`
}

func fltDotnetConsoleScript() string {
	return "echo \"not implemented\""
}
