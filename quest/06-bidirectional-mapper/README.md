## Quest 06: Bidirectional Mapper
### Task
You are a linguist in the Royal Canterlot Library, recently started on researching Prench the language. While most words are easy to understand enough, you still have difficulty in remembering certain obscure words. Moreover, you wanted to develop a tool to help other ponies also begin learning Prench, so you decided to write your own dictionary program able to find in both ways.

### Details
* The program will read from a dictionary file, which encodes maps in a certain format.
* The program will ask the user for the word they want to search. It can be either Ponish or Prench, and conjugations are not considered.
* Whatever the result is, it must be shown in the console.
* Example dictionary:
```
go,aller
learn,apprendre
pony,poney
love,aimer
apple,pomme
heart,coeur
horse,cheval
```
* Example output:
```
Search for word: apple
[Ponish] apple [Prench] pomme

Search for word: poney
[Prench] poney [Ponish] pony

Search for word: greed
Word "greed" could not be found.
```

### Requirements
* Name of the dictionary file can be hardcoded.
* Absolute file paths are prohibited.