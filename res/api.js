GET_my_collection = {
    code: int,
    message: String,
    data: [
        {
            id: String,
            name: String,
            url: String,
            image: String,
        },
    ]
}

POST_my_collection = {
    name: String,
    url: String,
    image: String,
}

POST_share_my_collection = {
    name: String
}

GET_collections = {
    code: int,
    message: String,
    data: [
        {
            id: String,
            name: String,
            owner: String,
        },
    ]
}

GET_collection_id = {
    code: int,
    message: String,
    data: [
        {
            id: String,
            name: String,
            url: String,
            image: String,
        },
    ]
}

PUT_take_collection_id = {}