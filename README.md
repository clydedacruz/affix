This mini project is meant to be a super simple blog or personal site for a single user.
From a readers perspective it is just a bunch of pages , accessible from an index page and a fast search.

from the  creater's perpective it has a clean markdown editor for creating or editing and deleting content.


technologies stack

bootstrap 
go gin-gonic
badger (key val store)
markdown pages 


# content

- content
  - my trip to krypton
    - 1.md
    - 2.md
  - bruce's party
    - 1.md
## A page. 
Each page is will be stored as a folder containing all the markdown for the page

## Index 
The index will be automatically generated from the list of folders in the content directory 


APIs 
GET  affix/topics/{topicid}/sheets/{sheetid}
POST affix/topics
{
  title
  metadata
  sheetPreview []byte
  sheets [][]byte
}
POST affix/topics/{topicid}/sheets
POST affix/topics/{topicid}/sheets/{sheetid}
PUT  affix/topics/{topicid}/sheets/{sheetid}
PUT affix/topics/{topicid}/sheets
