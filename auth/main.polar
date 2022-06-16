actor User{}

resource Book {
    roles = ["reader", "member", "librarian"];
    permissions =  ["read", "checkout", "checkin", "remove"];

    "read" if "reader";
    "checkout" if "member";
    "checkin" if "librarian";
    "remove" if "librarian";

    "member" if "librarian";
    "reader" if "member";
}

allow(actor, action, resource) if
    has_permission(actor, action, resource);
