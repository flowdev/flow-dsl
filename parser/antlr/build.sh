#!/bin/sh

echo "This script should be called from the dir where the script is stored!"

rm ./generated/*
java -jar ./antlr4-complete.jar -Dlanguage=Go -no-listener -o ./generated ./flowDslLexer.g4
java -jar ./antlr4-complete.jar -Dlanguage=Go -no-listener -o ./generated ./flowDsl.g4
cp ./generated/*.go ..
