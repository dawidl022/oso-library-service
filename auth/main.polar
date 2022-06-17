actor User {}

resource Book {
    roles = ["reader", "member", "librarian"];
    permissions = ["read", "checkout", "checkin", "remove"];

    "read" if "reader";
    "checkout" if "member";
    "checkin" if "librarian";
    "remove" if "librarian";

    "member" if "librarian";
    "reader" if "member";
}

has_role(actor: User, role_name: String, book: Book) if
    role_name = actor.Role and (
        book.GloballyAvailable or (reg in book.Regions and reg in actor.Regions)
    );

allow(actor, action, resource) if
    has_permission(actor, action, resource);
