gen_folder=./tower-sdk-go
src_folder=./tower-sdk-gencode

echo "|> Cleaning previously generated SDK"
if [ -d "$gen_folder" ]
then
    rm -rf $gen_folder/*
else
    mkdir $gen_folder
fi
rm -rf $src_folder/build

echo "|> Generating SDK"
cd $src_folder
./gradlew buildGoClient
cd -

echo "|> Moving generated files to '$gen_folder' folder"
mv $src_folder/build/go/* ./$gen_folder

echo "|> Patching the generated code package name"
sed -i 's#module github.com/GIT_USER_ID/GIT_REPO_ID#module openapi#g' $gen_folder/go.mod

echo "|> Patching the generated code, use local 'optional' module"
# replace (<spaces>*)Body optional.Map[string]interface{}(<spaces>*) with (<spaces>*)Body optional.Map(<spaces>*)
# go workspace automatically uses local 'optional' module
sed -i 's#^\(\s*\)Body optional.Map\[string\]interface{}\(\s*\)$#\1Body optional.Map\1#g' tower-sdk-go/api_default.go

