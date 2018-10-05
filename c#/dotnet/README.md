# HTTP serverを立ち上げる
```
dotnet new web
dotnet restore
dotnet run
```

# CUIを作る

## プロジェクトを作る
```
dotnet new console
```

## とりあえず走らせる
```
dotnet restore
dotnet run
```

## バイナリを作る
```
# csprojに RuntimeIdentifiersを追加
dotnet publish -c Release -f netcoreapp2.1 -r osx-x64
```
