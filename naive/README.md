# naive 

This more or less demonstrates the SDK's current module and API separation (there isn't one) what it might look like 
if separate go modules with SIV were introduced without any special consideration.

In a "real world" version of this scenario the folders {v1,v2} would be separate releases for each module as 
specified in the go.mod file.  This could only be demonstrated as changing over time or different 
branches in a single repo.  Different SIV versions of the same module are therefore demonstrated by changing the 
version in the go.mod file and duplicating all code to a separate folder.  This is not a realistic scenario, but it 
is the only way to demonstrate the problem in a single repo.  
