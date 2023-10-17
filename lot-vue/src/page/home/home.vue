<template>
  <div class="page">
    <div class="equipment">
      <label class="title">我的设备</label>
      <div class="equipmen-wrap">
        <div class="equipme-item" :style="{ minHeight: height * 0.1 + 'px' }">
          <div class="t">
            <div>插座</div>
            <div   v-if=" msg != '上线' ">离线</div>
            <img
              style="
                width: 38px;
                height: 26px;
                cursor: pointer;
                -webkit-tap-highlight-color: rgba(0, 0, 0, 0);
                -webkit-tap-highlight-color: transparent;
                -moz-tap-highlight-color: transparent;
              "
              :src="closeOpenIcon"
              v-else
              @click.prevent="changeStatus"
              alt=""
            />
           
          </div>
          <div class="t" style="align-items: baseline">
            <img
              style="
                cursor: pointer;
                margin-top: 0.3rem;
                -webkit-tap-highlight-color: rgba(0, 0, 0, 0);
                -webkit-tap-highlight-color: transparent;
                -moz-tap-highlight-color: transparent;
              "
              :style="{ height: height * 0.1 * 0.5 + 'px' }"
              src="../../images/设备_插座 圆 线.png"
              alt=""
            />
            
          </div>
        </div>
      </div>
      
    </div>
  </div>
</template>

<script>
import head from "../head/head";
import openIcon from "../../images/open.png";
import closeIcon from "../../images/close.png";
import fetch from "../../config/fetch";
import PahoMQTT from "paho-mqtt";
import { getStore, setStore } from '../../config/mUtils';
export default {
  data() {
    return {
      height: "",
      reconnectTimer:"",
      closeOpenIcon: closeIcon,
      open: false,
      msg: "",
      client: "",
      
    };
  },
  watch: {
            msg(newVal, oldVal) {
            // 监听 num 属性的数据变化
    		// 作用 : 只要 num 的值发生变化,这个方法就会被调用 
            //只要没有发生变化，就没有办法进行其他的操作
    		//newData是更新后的数据
            //oldData是旧数据
                // console.log('oldVal:',oldVal)
                // console.log('newVal:',newVal)
                this.msg = newVal
                // console.log("数据变化:",   this.msg);
            }
        }
,
  // mqtt: {
  //       'pp' (data) {
  //     this.msg = this.msg + data + '<br>'
  //     console.log(""+data)
  //   }
  // },
  mounted() {
    this.height = window.screen.height;
    this.init();
    console.log("status:",this.msg)
    // this.imitate()
    // console.log("容器高度",this.$refs.dataListContainer.$el.offsetHeight)
    // this.scroll();
    // this.initScrollChange()
  },

  computed: {},

  methods: {
   
    init() {
      if (this.open) {
        this.closeOpenIcon = openIcon;
      } else {
        this.closeOpenIcon = closeIcon;
      }
      this.initMqtt();
      //获取token
      fetch('/token').then(res=>{
        setStore('auth',res.data)
      })
    },

    initMqtt() {
      let that = this;
      let uniqueId = navigator.userAgent.replace(/[^\w]/gi, '');
      const mqttClient = new Paho.MQTT.Client("192.168.1.169", 9002, uniqueId);
      let i = 0;
      // 创建接选项
      const options = {
        timeout: 3,
        useSSL: false,
        onSuccess: onConnect,
        onFailure: onFailure,
      };

      // 连接成功回调函数
      function onConnect() {
        console.log("连接成功");
        // 订阅主题
        mqttClient.subscribe("lw");
        clearInterval(this.reconnectTimer)
        this.open = true
      }

      // 连接失败回调函数
      function onFailure() {
        console.log("连接失败");
        setInterval(function(){
           	mqttClient.connect(options);
           },3000)
      }

      // 接收消息回调函数
      function onMessageReceived(message) {
        // let data =  message.payloadString
        // that.msg = JSON.parse(data).msg
        // console.log(JSON.parse(data).msg)
      that.msg = message.payloadString
        // console.log("消息回掉",that.msg )
        if(message.payloadString == "上线"){
          that.open = true
        }else{
          that.open = false
        }

      }

      // 设置消息接收回调函数
      mqttClient.onMessageArrived = onMessageReceived;

      // 连接到MQTT代理
      mqttClient.connect(options);
      mqttClient.onConnectionLost = function (responseObject) {
        if(responseObject.errorCode !== 0) {
           alert("连接已断开");
           this.open = false
           this.reconnectTimer = setInterval(function(){
           	mqttClient.connect({
                cleanSession : true, 
                onSuccess : onConnect, 
                onFailure : onFailure, 
                keepAliveInterval: 3, 
                       // Enable automatic reconnect
            });
           },2000)
        }
    }
    },

    
    changeStatus() {
   
      
 
      if (!this.open) {
     
        fetch("/equipments/control", {
          order: "on",
        }).then((res) => {
          // console.log(res.data.msg);
          if (res.msg == "ok") {
            this.closeOpenIcon = openIcon;
            this.open = true;
          
          }
        });
      } else {
        fetch("/equipments/control", {
          order: "off  ",
        }).then((res) => {
          // console.log(res.data.msg);
          if (res.msg == "ok") {
            this.closeOpenIcon = closeIcon;
            this.open = false;
          }
        });
      }
    },
  },
  components: {
    Head: head,
  },
  watch: {
    inputValue(newName, oldName) {
      this.watchInputValue = newName;
    },
  },
};
</script>

<style lang="scss" scoped>
body {
  -webkit-tap-highlight-color: rgba(0, 0, 0, 0);
  -webkit-tap-highlight-color: transparent;
  -moz-tap-highlight-color: transparent;
}

.page {
  /* 直接在style标签中设置背景图片 */
  background: #f4f4f4;
  height: 100vh;

  /* 添加其他样式，确保页面内容可见 */
  color: #333; /* 文字颜色，以便在背景图片上可见 */
  /* 其他样式 */
  .equipment {
    width: 94%;
    margin: 0 auto;
    .title {
      display: block;
      padding: 10px 0;
      font-size: 20px;
      font-family: 黑体;
    }
    .equipmen-wrap {
      margin-top: 20px;
      display: flex;
      flex-flow: wrap;
      justify-content: space-between;
      .equipme-item {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        width: 44%;
        border-radius: 6px;
        background: #ffffff;
        color: #333;
        padding: 3%;
        .t {
          display: flex;
          justify-content: space-between;
          .manage-btn {
            font-size: 14;
          }
        }
      }
    }
  }
}
</style>
