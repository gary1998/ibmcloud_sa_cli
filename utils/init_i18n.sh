#!/bin/bash
echo String collection file: $1
echo Resource file to be created: $2
echo Selected i18n dir: $3

echo Attempting German translation
python3 translator.py $1 en-de $3/en_ES.all.json && echo German translation completed successfully! || exit 0

echo Attempting Spanish translation
python3 translator.py $1 en-es $3/en_ES.all.json && echo Spanish translation completed successfully! || exit 0

echo Attempting French translation
python3 translator.py $1 en-fr $3/en_ES.all.json && echo French translation completed successfully! || exit 0

echo Attempting Italian translation
python3 translator.py $1 en-it $3/en_ES.all.json && echo Italian translation completed successfully! || exit 0

echo Attempting Japanese translation
python3 translator.py $1 en-ja $3/en_ES.all.json && echo Japanese translation completed successfully! || exit 0

echo Attempting Korean translation
python3 translator.py $1 en-ko $3/en_ES.all.json && echo Korean translation completed successfully! || exit 0

echo Attempting Portuguese translation
python3 translator.py $1 en-pt $3/en_ES.all.json && echo Portuguese translation completed successfully! || exit 0

echo Attempting Chinese (Simplified) translation
python3 translator.py $1 en-zh $3/en_ES.all.json && echo Chinese (Simplified) translation completed successfully! || exit 0

echo Attempting Chinese (Traditional) translation
python3 translator.py $1 en-zh-TW $3/en_ES.all.json && echo Chinese (Traditional) translation completed successfully! || exit 0

go-bindata -pkg resources -o $1 $2 && echo Translations saved successfully!