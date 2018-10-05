OS = osx-x64
APPVERSION = 2.1

.PHONY: run build
run:
	dotnet run
build:
	dotnet publish -c Release -f netcoreapp${APPVERSION} -r ${OS}
