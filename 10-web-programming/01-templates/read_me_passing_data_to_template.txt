you can pass your own functions / predefined global functions thats alread available for use inside your templates.
To pass functions into our template, we need to have some data structure that aggregates together the functions we want to pass - that data structure is a map. - there is a special type in text/template package called as 'funcMap'
map of string and (empty)interface

An Empty interface is a interface without values, every type has atleast no methods. So every type implements the empty interface.
