mongosh -- "$MONGO_INITDB_DATABASE" <<EOF
db = db.getSiblingDB('$MONGO_INITDB_DATABASE')
db.auth('$MONGO_USERNAME', '$MONGO_PASSWORD');

// ----------------------------
// Collection structure for categories
// ----------------------------
db.getCollection("categories").drop();
db.createCollection("categories");
db.categories.createIndex({ "name": -1 });



// ----------------------------
// Collection structure for comment
// ----------------------------
db.getCollection("comment").drop();
db.createCollection("comment");
db.comment.createIndex({ "post_info.post_id": 1 });
db.comment.createIndex({ "create_time": -1 });

// ----------------------------
// Collection structure for configs
// ----------------------------
db.getCollection("configs").drop();
db.createCollection("configs");

db.configs.createIndex({ "typ": 1 });

// ----------------------------
// Documents of configs
// ----------------------------
db.getCollection("configs").insert([ {
    _id: "webmaster",
    typ: "webmaster",
    props: {
        name: "fnote",
        postCount: 0,
        categoryCount: 0,
        websiteViews: 0,
        websiteLiveTime: Date.now(),
        profile: "hello, fnote",
        picture: "",
        websiteIcon: "",
        domain: ""
    },
    create_time: Date.now(),
    update_time: Date.now()
} ]);
db.getCollection("configs").insert([ {
    _id: "comment",
    typ: "comment",
    props: {
        status: true
    },
    create_time: Date.now(),
    update_time: Date.now()
} ]);
db.getCollection("configs").insert([ {
    _id: "friend",
    typ: "friend",
    props: {
        status: false
    },
    create_time: Date.now(),
    update_time: Date.now()
} ]);


// ----------------------------
// Collection structure for friends
// ----------------------------
db.getCollection("friends").drop();
db.createCollection("friends");
db.getCollection("friends").createIndex({
    url: NumberInt("1")
}, {
    name: "url_1",
    unique: true
});
// 创建 create_time 降序索引
db.friends.createIndex({ "create_time": -1 });


// ----------------------------
// Collection structure for message_template
// ----------------------------
db.getCollection("message_template").drop();
db.createCollection("message_template");
// 创建 name 升序索引
db.message_template.createIndex({ "name": 1 });

// ----------------------------
// Documents of message_template
// ----------------------------
db.getCollection("message_template").insert([ {
    _id: "friend",
    name: "friend",
    title: "友链申请通知",
    content: "您好，您的网站有了新的友链申请，详情可前往后台查看。",
    create_time: Date.now(),
    update_time: Date.now(),
    active: 1,
    recipient_type: 0
} ]);
db.getCollection("message_template").insert([ {
    _id: "comment",
    name: "comment",
    title: "文章评论通知",
    content: "您好，您在文章有新的评论，详情请前往后台进行查看。",
    create_time: Date.now(),
    update_time: Date.now(),
    recipient_type: 0,
    active: 1
} ]);

// ----------------------------
// Collection structure for posts
// ----------------------------
db.getCollection("posts").drop();
db.createCollection("posts");
// 创建 create_time 降序索引
db.posts.createIndex({ "create_time": -1 });
// 创建 create_time 升序索引
db.posts.createIndex({ "create_time": 1 });
// 创建 category 单字段索引
db.posts.createIndex({ "category": 1 });
// 创建 category 和 tags 复合索引
db.posts.createIndex({ "category": 1, "tags": 1 });
// 创建 title 文本索引
db.posts.createIndex({ "title": "text" });

// ----------------------------
// Collection structure for visit_logs
// ----------------------------
db.getCollection("visit_logs").drop();
db.createCollection("visit_logs");
// 创建 create_time 降序索引
db.visit_logs.createIndex({ "create_time": -1 });
EOF