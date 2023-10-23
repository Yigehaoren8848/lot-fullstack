<template>
    <div class="bottom-navbar">
      <router-link to="/home">首页</router-link>
      <router-link to="/mine">我的</router-link>
      
    </div>
  </template>
  
  <script>
  export default {
    
  };
  </script>
  
  <style scoped>
  a{
    color: white;
  }
  /* 样式可以根据需求自定义 */
  .bottom-navbar {
    position: fixed;
    bottom: 0;
    left: 0;
    width: 100%;
    height: 40px;
    line-height: 40px;
    font-size: 20px;
    display: flex;
    justify-content: space-around;
    background-color: #333;
  
    padding: 10px 0;
  }
  
  .router-link-active {
    font-weight: bold;
  }
  .router-link-active{
    font-weight: normal;
  }
  </style>
  