# How to get GAS_KEY

GAS is Google Apps Script.


## Step.1

**Access** [Google Drive](https://drive.google.com/drive) and **Sign Up/In**.


## Step.2

**Press** New button in upleft, and **Select** Google Apps Script.


## Step.3

**Copy and Paste** source code in below.

```
function doPost(e) {
  var p = e.parameter;
  var translatedText = LanguageApp.translate(p.text, p.source, p.target);
  return ContentService.createTextOutput(translatedText);
}
```


## Step.4

**Deploy** Google Apps Script


## Step.5

**Copy** deploy ID, and **Set** environment variable.
