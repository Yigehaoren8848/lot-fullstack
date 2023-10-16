<template>
    <div style="margin: 1rem 0;">
        <div class="loading"></div>



        <slot></slot>
    </div>
</template>

<script>

export default{
    data() {
    return {
        categoryArrData:[],
        showMore:true
    };
  },

  props: {
    showLoading:{
        type:Boolean,
        default:false
    }
  },
  mounted() {
  
  },

  components: {},

  computed:{
   
  },

  methods: {
   
    reload() {
      window.location.reload();
    },
  },
};

</script>

<style lang="scss" scoped>
.loading {
  display: block;
  position: relative;
  width: 4px;
  height: 5px;
  margin: 0 auto;
  margin-top: .1rem;
  animation: rectangle infinite 1s ease-in-out -0.2s;
  background-color: #409eff;
}

.loading:before,
.loading:after {
  position: absolute;
  width: 4px;
  height: 5px;
  content: "";
  background-color: #409eff;
}

.loading:before {
  left: -7px;

  animation: rectangle infinite 1s ease-in-out -0.4s;
}

.loading:after {
  right: -7px;

  animation: rectangle infinite 1s ease-in-out;
}

@keyframes rectangle {
  0%,
  80%,
  100% {
    height: 10px;
    box-shadow: 0 0 #409eff;
  }

  40% {
    height: 15px;
    box-shadow: 0 -10px #409eff;
  }
}
</style>