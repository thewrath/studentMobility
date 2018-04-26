Command line to test server :
to test put student Data : 
	curl -s -vX  PUT -d @createStudent.json http://localhost:3000/putStudentData
to test getStudent Data : 
	curl -s -vX  GET -d @getStudent.json http://localhost:3000/getStudentData
to test removeStudentData : 
	curl -s -vX  POST -d @removeStudent.json http://localhost:3000/removeStudentData

You can find inside student.json and year.json how correctly form a json for store student data. 

So send this json to the server with the folowing http : 

getStudentData (response with a json with student data inside)
putStudentData (response with OK if your request is correctly handle by the server, if your student data don't have id value server create a new student inside DB)
removeStudentData (remove all student data of the given student id, to remove only fields of student send a putStudentData with fields, that you want to remove, empties)


