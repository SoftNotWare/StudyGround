document.addEventListener('DOMContentLoaded', function (event) {
    ABP.create(document.getElementById("load-player"), {
        "src": {
            "playlist": [
                {
                    "video": document.getElementById("video-1"),
                    "comments": "danmu/comment-otsukimi.xml"
                },
                {
                    "video": document.getElementById("video-2"),
                    "comments": "danmu/comment-science.xml"
                }
            ]
        },
        "width": 1000,
        "height": 690
    })
})

new Vue({
    el: '#play',
    data: {
        playlist: [
            { 'id': 1, 'title': '两只黄鹂鸣翠柳', 'active': true },
            { 'id': 2, 'title': '一行白鹭上青天', 'active': false },
            { 'id': 3, 'title': '窗含西岭千秋雪', 'active': false },
            { 'id': 4, 'title': '门泊东吴万里船', 'active': false },
        ],
        danmulist: [
            { 'id': 1, 'title': '两只黄鹂鸣翠柳', 'active': false },
            { 'id': 2, 'title': '一行白鹭上青天', 'active': false },
            { 'id': 3, 'title': '窗含西岭千秋雪', 'active': false },
            { 'id': 4, 'title': '门泊东吴万里船', 'active': false },
        ]
    }
})