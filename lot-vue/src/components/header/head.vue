<template>
    <header id='head_top'>
        <div v-show="showMenu" class="main-menu">
            <div class="main-menu-item">发布</div>
            <router-link :to="ucenter" class="main-menu-item" style="display: block;color: white;">个人中心</router-link>
            <div class="main-menu-item">登录</div>
       
        </div>
        <slot name='logo'></slot>
        <router-link
        
        :to="home"
        :tag="div"
        style="line-height: 2.8rem; color: white; margin-left: 2%"
      >
        多线程
      </router-link>
        <div @click.prevent="changeMenu" style="width: 1.8rem;margin-right: 2%;" :to="userInfo? '/profile':'/login'" v-if='signinUp' class="head_login">
            <img src="../../images/menu.png" style="height: 80%;margin-top: .3rem;" alt="">

        </div>
        <section class="title_head ellipsis" v-if="headTitle">
            <span class="title_text">{{headTitle}}</span>
        </section>
    </header>
</template>

<script>
    import {mapState, mapActions} from 'vuex'
    export default {
    	data(){
            return{
                showMenu:false,
                home:'/home',
                ucenter:'/ucenter'
            }
        },
        mounted(){
            //获取用户信息
            this.getUserInfo();

        },
        props: ['signinUp', 'headTitle', 'goBack'],
        computed: {
            ...mapState([
                'userInfo'
            ]),
        },
        methods: {
            changeMenu(){
                this.showMenu = !this.showMenu
            }
        },

    }

</script>

<style lang="scss" scoped>
    @import '../../style/mixin';
.main-menu{
    position: fixed;

  left: 50%;
  transform: translate(-50%, 0%);
  width: 100%;
  background-color: #409eff;
  margin-top: 2.8rem;
  color: white;
  margin-bottom: 6rem;

}
.main-menu-item{
    text-align: center;  padding: .3rem 0;
}
    #head_top{
        background-color: #409eff;
        position: fixed;
        display: flex;
        justify-content: space-between;
        z-index: 100;
        left: 0;
        top: 0;
        @include wh(100%, 2.8rem);
    }
    .head_goback{
        left: 0.4rem;
        @include wh(0.6rem, 1rem);
        line-height: 2.2rem;
        margin-left: .4rem;
    }
    .head_login{
     
   
    }
    .title_head{
        @include center;
        width: 50%;
        color: #fff;
        text-align: center;
        .title_text{
            @include sc(0.8rem, #fff);
            text-align: center;
            font-weight: bold;
        }
    }
</style>
