- Use fmt.Println to see output in the terminal to easily debug/logging.
- Use fmt.Fprintf to send output to the user/browser.

- in my struct model id is int then why I'm idStr := r.PathValue("itemId") doing this to collect id as string in idStr and why later using strconv.Atoi(idStr) to make it int again?
  Ans: All values extracted from an HTTP request—such as path parameters, query parameters, and form values—are always received as strings, because HTTP is a text-based protocol. And to match with the database id type we need to convert the url id in integer again then compare it or by using it we'll try to find specific data in db.

- func Use(middlewares ...Middleware)
  Here, ... is after the type (Middleware) → variadic parameter
  Means: “this function can accept any number of Middleware arguments”
  Inside the function, it behaves like a slice: []Middleware

- append(mngr.globalMiddlewares, middlewares...)
  Here, ... is after a slice variable → “expand the slice” or spreading
  Means: “take all elements of this slice and pass them individually”
  Needed because append expects multiple elements, not a slice

- Environment variables are always strings that's why I need to convert those in other format if i need.
