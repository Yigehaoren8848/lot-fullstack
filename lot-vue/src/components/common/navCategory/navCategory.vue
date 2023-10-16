<template>
  <div>
    <div class="catetgory-container">
      <div
        class="catetgory-wrap"
        v-for="(item, index) in categoryArrData"
        :key="index"
      >
        <div class="category-name">{{ item.categoryName }}</div>
        <div
          v-for="(item, index) in item.categoryDetail"
          :key="index"
          class="category-arr"
        >
          {{ item }}
        </div>
        
      </div>
      <div v-show="showMore" class="more" @click.prevent="more">更多</div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
        categoryArrData:[],
        showMore:true
    };
  },

  props: {
    categoryArr: {
      type: Array,
      default: () => [
        { categoryName: "类型", categoryDetail: ["大姐", "二姐", "三姐"] },
        { categoryName: "当地", categoryDetail: ["大姐", "二姐", "三姐"] },
      ],
    },
  },
  mounted() {
    this.categoryArrData = this.categoryArr
  },

  components: {},

  computed:{
   
  },

  methods: {
    more(){
        let more = [
        { categoryName: "订单", categoryDetail: ["多少度", "撒似懂非懂", "sdd"] },
        { categoryName: "没得", categoryDetail: ["ssd", "dsf", "第三方"] },
      ]
      this.categoryArrData = this.categoryArrData.concat(more)
      this.showMore = false
    },
    reload() {
      window.location.reload();
    },
  },
};
</script>


<style lang="scss" scoped>
.catetgory-container {
  background: white;
  border-radius: 0.2rem;
  padding: 0.3rem;
  margin: 0 auto;
  margin-top: 2.7rem;
  width: 96%;
}

.catetgory-wrap {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
}
.category-name {
}
.category-arr {
  padding-left: 0.3rem;
  color: #4f4f44;
}
.category-arr:hover {
  color: #409eff;
  cursor: pointer;
}
.more{
    display:flex;justify-content: end;padding-right: .5rem;color: #4f4f4f;cursor: pointer;
}
</style>