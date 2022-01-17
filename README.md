# Kubectl ServiceAccount Impersonation

The projet name `assa` is first and foremost a palindrome, which is always fun, but it also means `as sa`.

It allows you to run kubectl commands as a ServiceAccont.

Kubectl lets you impersonate a user or a group, but not a service account. This may seem surprising
at first because users, groups and service accounts are all equals when it comes to role bindings.

However, you have to remember that kubectl is a client-side tool. When it sends requests to the Kubernetes
API server it needs to send the identity of the requester. This identity can be only a user or a group.





