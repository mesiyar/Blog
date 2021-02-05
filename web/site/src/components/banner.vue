<template>
    <div id="banner" :class="{'home-banner':isHome}">
        <div class="banner-img" :style="{'background-image': `url(${src})`}">
            <template v-if="isHome">
                <!--博主信息-->
                <div class="focusinfo">
                    <!-- 头像 -->
                    <div class="header-tou">
                        <router-link to="/"><img :src="websiteInfo.avatar"></router-link>
                    </div>
                    <!-- 简介 -->
                    <div class="header-info">
                        <p>{{websiteInfo.slogan}}</p>
                    </div>
                    <!-- 社交信息 -->
                    <div class="top-social">
                        <div v-for="item in socials" :key="item.id" :title="item.title"><a :href="item.href" target="_blank" :style="{'color':item.color}"><i class="iconfont" :class="item.icon"></i></a></div>
                    </div>
                </div>
                <!--左右倾斜-->
                <div class="slant-left"></div>
                <div class="slant-right"></div>
            </template>
        </div>
    </div>
</template>

<script>
    import {fetchSiteInfo} from '../api'
    export default {
        name: "banner",
        data(){
            return{
                websiteInfo: {
                    avatar: 'https://s2.ax1x.com/2020/01/17/1SCadg.png',
                    slogan: 'The way up is not crowded, and most chose ease.',
                    site_name: 'LearnAndGo',
                    domain: 'http://www.learnkai.top',
                    desc: '一个It技术的追随者'
                },
                socials: [
                  {
                        id: 1,
                        title: 'Gitee',
                        icon: 'icongitee',
                        color: '#d81e06',
                        href: 'https://gitee.com/mesiyar'
                    },
                    {
                        id: 2,
                        title: 'GitHub',
                        icon: 'icongithub',
                        color: '',
                        href: 'https://github.com/mesiyar'
                    },
                ]
            }
        },
        props:{
            src:{
                type: String,
                default: 'http://api.learnkai.top/upload/images/bg.jpg'
            },
            isHome:{
                type: [Boolean,String],
                default: false
            }
        },
        created(){
           this.fetchSiteInfo()
        },
        methods:{
          fetchSiteInfo() {
                fetchSiteInfo().then(res => {
                    this.websiteInfo = res.data
                    console.log(res)
                }).catch(err => {
                    console.log(err)
                })
            }
        }
    }
</script>

<style scoped lang="less">

    #banner {
        position: relative;
        margin-top: 80px;
        width: 100%;
        height: 500px;
        .banner-img{
            width: inherit;
            height: inherit;
            background-position: center;
            background-size: cover;
            background-repeat: no-repeat;
            transition: all 0.2s linear;
            overflow: hidden;
            &:hover {
                transform: scale(1.1, 1.1);
                filter: contrast(130%);
            }
        }
        &.home-banner {
            height: 550px;
            .banner-img{
                background-position: center center;
                background-repeat: no-repeat;
                background-attachment: fixed;
                background-size: cover;
                z-index: -1;
                transition: unset;
                &:hover {
                    transform: unset;
                    filter: unset;
                }
            }
            .slant-left {
                content: '';
                position: absolute;
                width: 0;
                height: 0;
                border-bottom: 100px solid #FFF;
                border-right: 800px solid transparent;
                left: 0;
                bottom: 0;
            }
            .slant-right {
                content: '';
                position: absolute;
                width: 0;
                height: 0;
                border-bottom: 100px solid #FFF;
                border-left: 800px solid transparent;
                right: 0;
                bottom: 0;
            }
        }
    }
    .focusinfo {
        position: relative;
        max-width: 800px;
        padding: 0 10px;
        top: 40%;
        left: 50%;
        transform: translate(-50%,-50%);
        -webkit-transform: translate(-50%,-50%);
        text-align: center;
        img {
            width: 80px;
            height: auto;
            border-radius: 50%;
            border: 3px solid rgba(255, 255, 255, 0.3);
        }
        .header-info {
            width: 60%;
            font-size: 14px;
            color: #EAEADF;
            background: rgba(0, 0, 0, 0.66);
            padding: 20px 30px;
            margin: 30px auto 0 auto;
            font-family: miranafont,"Hiragino Sans GB",STXihei,"Microsoft YaHei",SimSun,sans-serif;
            letter-spacing: 1px;
            line-height: 30px;
        }
        .top-social {
            height: 32px;
            margin-top: 30px;
            margin-left: 10px;
            list-style: none;
            display: inline-block;
            font-family: miranafont,"Hiragino Sans GB",STXihei,"Microsoft YaHei",SimSun,sans-serif;
            div {
                float: left;
                margin-right: 10px;
                height: 32px;
                line-height: 32px;
                text-align: center;
                width: 32px;
                background: white;
                border-radius: 100%;
            }
        }
    }
    @media (max-width: 960px){
        #banner {height: 400px;}
    }
    @media (max-width: 800px){
        #banner {display: none;}
    }
</style>
