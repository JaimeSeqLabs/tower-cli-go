echo "|> Cleaning previously generated SDK"
rm -rf gen/*
rm -rf tower-sdk-gencode/build

echo "|> Generating SDK"
cd tower-sdk-gencode
./gradlew buildGoClient
cd -

echo "|> Moving generated files to 'gen' folder"
mv tower-sdk-gencode/build/go/* ./gen

echo "|> Patching the generated code package name"
sed -i 's#module github.com/GIT_USER_ID/GIT_REPO_ID#module openapi#g' gen/go.mod

