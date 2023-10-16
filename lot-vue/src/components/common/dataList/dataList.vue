<template>
  <div
  ref="scroll"
    
    id="box"
    @touchend="touchend"
    @touchmove="touchmove"
    @touchstart="touchstart"
    class="data-list-container"
  >
    <router-link
    to="/project_detail"
      ref="ds"
      class="data-list-item"
      :class="{ alignCenter: showImg }"
      :style="{ 'background-color': bg, width: width, height: '6.5rem' }"
      v-for="(item, index) in dataList"
      :key="index"
    >
      <div class="img-container" :style="{ height: '6rem' }" v-if="showImg">
        <img
          class="data-list-item-img"
          src="https://img1.baidu.com/it/u=3778268823,1285877823&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1685811600&t=a732fa0319a856c8f9a0c7b3b125cd5c"
          alt=""
        />
      </div>
      <div class="text-container">
        <div class="main-content">{{ item }}</div>
        <div class="other-contnt">
          <div style="color: #333">金额</div>
          <div style="color: red; margin-left: 0.6rem">340</div>
        </div>
        <div class="other-contnt">
          <div style="color: #333">发布时间</div>
          <div style="color: red; margin-left: 0.6rem">340</div>
        </div>
        <div class="other-contnt">
          <div style="color: #333">浏览量</div>
          <div style="color: red; margin-left: 0.2rem">340</div>
          <div style="color: red; margin-left: 0.4rem; margin-right: 0.4rem">
            |
          </div>
          <div style="color: #333">浏览量</div>
          <div style="color: red; margin-left: 0.2rem">340</div>
        </div>
      </div>
    </router-link>
  </div>
</template>
<script>
export default {
  data() {
    return {
      touchEndTitleShow: false, //控制手指离开屏幕的title显示
      touchstartTitleShow: false, //控制手指按下屏幕的title显示
      number: 0, //列表回弹动画时间
      translateY: 0, //列表随手指下拉而偏移的量
      startY: 0, //手指按住的位置的y坐标，也就是起始坐标
      hasNext: true, //是否还有下一页
      loading: false, //loading显示
    };
  },

  props: {
    bg: {
      tyte: String,
      default: "red",
    },
    width: {
      type: String,
      default: "98%",
    },
    height: {
      type: String,
      default: "3rem",
    },
    dataList: {
      type: Array,
      default: () => [],
    },
    showImg: {
      type: Boolean,
      default: true,
    },
  },
  mounted() {
    //   console.log(this.$refs.ds)
    // this.initScrollChange()
  },

  components: {},

  computed: {
    //将获取的数据按照A-Z字母开头排序
    sortgroupcity() {
      let sortobj = {};
      for (let i = 65; i <= 90; i++) {
        if (this.groupcity[String.fromCharCode(i)]) {
          sortobj[String.fromCharCode(i)] =
            this.groupcity[String.fromCharCode(i)];
        }
      }
      return sortobj;
    },
  },

  methods: {
    //点击图标刷新页面
    reload() {
      window.location.reload();
    },
    //手指触碰到屏幕
    touchstart(e) {
       
        this.number = 0
        let y = e.targetTouches[0].pageY
        this.startY = y
      },
      //手指开始滑动
      touchmove(e) {
        
        let y = e.targetTouches[0].pageY
        if((y>this.startY)&&this.$refs.scroll.scrollTop==0){
          this.touchstartTitleShow = true
          //如果当前移动距离大于初始点击坐标，则视为是下拉。并且要处于顶部才刷新，不能影响正常的列表滑动。
          this.translateY = (y-this.startY)/2
        }
      },
      //手指松开
      touchend(e){
   
        let y = e.changedTouches[0].pageY
       if(y>this.startY){
         this.number = 4
         this.translateY = 0
         this.touchstartTitleShow = false
         this.touchEndTitleShow = true
         setTimeout(()=>{
           this.touchEndTitleShow = false
         },1000)
         this.startY = 0
       }
      },
      initScrollChange() {
       this.$refs.scroll.onscroll  = (e) => {
        
         const offsetHeight = this.$refs.scroll.offsetHeight //可视区域的高度
         const scrollHeight = this.$refs.scroll.scrollHeight //元素全部高度
         const scrollTop = this.$refs.scroll.scrollTop //滚动条滚动距离
         console.log("scroll事件，",offsetHeight,"，",scrollHeight,"，",scrollTop)
         //可视区域高度加上滚动条滚动距离大于等于元素全部高度则表示滚动到底
         if ((offsetHeight + scrollTop) - scrollHeight >= -1) {
           console.log('到底啦')
           if (!this.loading && this.hasNext) {
               
             this.getData()
             
           }
         }
       }
     },
     getData() {
       this.loading = true
       
       setTimeout(() => {
           
         for (let i = 0; i < 10; i++) {
           this.dataList.push(this.dataList.length + 1)

         }
         this.loading = false
         if (this.list.length > 30) {
           this.hasNext = false
         }
       }, 1000)
     }
  },

};
</script>

<style lang="scss" scoped>
.data-list-item {
  padding: 0.3rem;
  margin: 0.3rem auto;
  border-radius: 0.3rem;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
.img-container {
  width: 30%;
  // height: 2rem;
}
.data-list-item-img {
  width: 100%;
  height: 100%;
}
.text-container {
  width: 70%;
  margin-left: 1rem;
}
.alignCenter {
  align-items: center;
}
.main-content {
  padding-top: 0rem;
  height: 70%;
}
.other-contnt {
  font-size: 0.8rem;

  display: flex;
}
.associate-contnt {
  padding-top: 0rem;
  height: 30%;
}
</style>