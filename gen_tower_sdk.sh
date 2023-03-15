echo "|> Cleaning previously generated SDK"
rm -rf gen/*
rm -rf tower-sdk-gencode/build

echo "|> Generating SDK"
cd tower-sdk-gencode
./gradlew buildGoClient
cd -

echo "|> Moving generated files to 'gen' folder"
mv tower-sdk-gencode/build/go/* ./gen
