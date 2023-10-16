<!-- <template>
  <div style="margin-top: 2.8rem">
    <headTop headTitle="项目详情" signinUp="projectDetail">
      
    </headTop>
    <div class="wrap">
      <div class="projectName">梵蒂冈电饭锅梵蒂冈单独发啥说法搭嘎</div>
      <div class="projectTime f4color flex mt">
        <div>发布日期：</div>
        <div>2022-11-12</div>
      </div>
      <div class="price f4color flex mt">
        <div>价格：</div>
        <div>1000000</div>
      </div>
      <div class="projectViews f4color flex mt">
        <div>浏览量：</div>
        <div>10000</div>
      </div>
      <div class="projectViews f4color flex mt">
        <div>报名人数：</div>
        <div>10000</div>
      </div>
    </div>

    <div class="demand" style="margin-top: 0.4rem">
      <div class="projectName pd-bto">要求</div>
      <div class="projectContent">
        梵蒂冈电饭锅的 2.props校验
        组件可以指定props验证要求，这对开发第三方组件来说，可以让使用者更加准确地使用组件。使用验证的时候，props接收的参数为json对象，而不是上述例子中的数组，例如：props:{a:Number},即为验证参数a需要为Number类型，如果调用该组件传入的a参数为字符串则会抛出异常。
        Vue提供props验证的方式有很多中，具体如下：
        （1）基本类型检测：prop接收的参数为原生构造器，Stirng，Number，Boolean，Function，Object，Array。也可以接收null，意味任意类型均可。
        （2）多种类型：prop:[Number,String];允许参数为多种类型之一；
        （3）参数必须:prop:[type:Number,required:true];表示参数必须有值且为Number类型。
        （4）参数默认值：prop:[type:Number,default:10]
        表示参数具有默认值10；需要注意的是如默认值设置为数组或对象，需要向组件中data属性那样，通过函数返回值得形式复制，例如：
      </div>
    </div>
    <div class="file" style="margin-top: 0.4rem">
      <div class="pd-bto">解决方案</div>
      <div class="flex">
        <img
          style="width: 4rem; height: 4rem; padding-right: 1rem"
          src="../../../images/zip.png"
          alt=""
        />
        <img
          style="width: 4rem; height: 4rem; padding-right: 1rem"
          src="../../../images/word.png"
          alt=""
        />
      </div>
    </div>

    <div style="position: fixed; bottom: 0.6rem; right: 2%">
      <el-button type="primary">报名</el-button>
    </div>
  </div>
</template>

<script>
import headTop from "../../header/head";

export default {
  data() {
    return {
      categoryArrData: [],
      showMore: true,
    };
  },

  props: {
    projectId: {
      type: String,
      default: "",
    },
    itemsData: {
      type: Object,
      default: () => {
        return {};
      },
    },
  },
  mounted() {},

  components: {
    headTop,
  },

  computed: {},

  methods: {
    reload() {
      window.location.reload();
    },
  },
};
</script>

<style lang="scss" scoped>
.wrap {
  background: white;
  width: 96%;
  margin: 0 auto;
  // border-radius: .3rem;
  padding: 0.3rem;
}
.el-button--primary {
  width: 100px;
}
.projectContent {
  color: #5d5d5d;
}
.f4color {
  color: #5d5d5d;
  font-size: 0.8rem;
}
.flex {
  display: flex;
}
.mt {
  margin-top: 0.15rem;
}
.demand {
  min-height: 30vh;
  background-color: white;
  width: 96%;
  margin: 0 auto;
  // border-radius: .3rem;
  padding: 0.3rem;
}
.file {
  min-height: 12vh;
  background-color: white;
  width: 96%;
  margin: 0 auto;
  // border-radius: .3rem;
  padding: 0.3rem;
}
.pd-bto {
  padding-bottom: 0.3rem;
}
</style> -->