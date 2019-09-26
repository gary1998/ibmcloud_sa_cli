import json
import sys
import json
from ibm_watson import LanguageTranslatorV3

language_translator = LanguageTranslatorV3(
    version='2018-05-01',
    iam_apikey='d-RlCiIJFUaNgmVtghTn3CGnxvaX_vQo37e98t6N8Cgg',
    url='https://gateway-lon.watsonplatform.net/language-translator/api'
)

fileName = sys.argv[1]
method = sys.argv[2]
finalFile = sys.argv[3]

print("Translating strings in file: "+fileName+" using method: "+method)

finalData = []

with open(fileName, 'r') as fileData:
  jsonData = json.load(fileData)
  for i in range(0, len(jsonData)):
    msg = jsonData[i]["translation"]
    translation = language_translator.translate(
      text=msg,
      model_id=method).get_result()
    translatedData = translation["translations"][0]["translation"]
    jsonRecord = {
      'id' : msg,
      'translation' : translatedData,
      'modified' : 'false'
    }
    finalData.append(jsonRecord)

with open(finalFile, 'w') as fileData:
  fileData.write(json.dumps(finalData, ensure_ascii=False))

print("Translated successfully!")