

new Vue({
    el: '#app',
    data: {
        carousels: [
            { text: '看成败', color: 'primary' },
            { text: '人生豪迈', color: 'info' },
            { text: '只不过是', color: 'success' },
            { text: '从头再来', color: 'warning' },
            { text: '……', color: 'danger' }
        ],
        activeTab: 0,
        tabs: [
            {
                label: '独立游戏',
                icon: 'google-photos'
            },
            {
                label: '游戏设计',
                icon: 'library-music'
            },
            {
                label: '职场技能',
                icon: 'video'
            }
        ],
        videoPage: 0,
        tab1: [
            {
                title: '母猪的产后护理',
                image: 'image/placeholder-1280x960.png',
                author: '大卫路易斯'
            },
            {
                title: '仙人掌嫁接技术',
                image: 'image/placeholder-1280x960.png',
                author: '卢比索亚森'
            },
            {
                title: '喵星人的自我修养',
                image: 'image/placeholder-1280x960.png',
                author: '杰尼狄娜'
            },
            {
                title: '不是所有汪都是乖宝',
                image: 'image/placeholder-1280x960.png',
                author: '徐甴莱'
            }
        ],
        tab2: [
            {
                title: '艾洛诗集',
                image: 'image/placeholder-1280x960.png',
                author: '这顿我请'
            },
            {
                title: '我为什么这么开',
                image: 'image/placeholder-1280x960.png',
                author: '瑾雅'
            },
            {
                title: '赵大脑袋战记',
                image: 'image/placeholder-1280x960.png',
                author: '赵大'
            },
            {
                title: '点亮你的猴生',
                image: 'image/placeholder-1280x960.png',
                author: '一不拉稀摸围棋'
            }
        ]
    },
    methods: {
        info(value) {
            this.videoPage = value
        }
    }
})