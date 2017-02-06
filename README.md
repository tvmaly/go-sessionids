# go-sessionids

a simple function to generate string based session ids for your web projects in Go

example:

   session_id :=  GenerateSessionId(40) // generate a random value 40 bytes in length

   this returns a base64 encoded session value that is 56 characters in length due to encoding.
