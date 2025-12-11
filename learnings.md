- Use fmt.Println to see output in the terminal to easily debug/logging.
- Use fmt.Fprintf to send output to the user/browser.

- in my struct model id is int then why I'm idStr := r.PathValue("itemId") doing this to collect id as string in idStr and why later using strconv.Atoi(idStr) to make it int again?
  Ans: All values extracted from an HTTP request—such as path parameters, query parameters, and form values—are always received as strings, because HTTP is a text-based protocol. And to match with the database id type we need to convert the url id in integer again then compare it or by using it we'll try to find specific data in db.
