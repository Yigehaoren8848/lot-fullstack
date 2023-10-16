<template>
    <div>
      <!-- Your component template here -->
      wgs:[{{ wgsLng }},{{ wgsLat }} ]
      gc:[{{ gcjLng }},{{ gcjLat }} ]
    </div>
  </template>
  
  <script>
  // 引入coordtransform库
  import coordtransform from 'coordtransform';
  
  export default {
    name: 'YourComponent',
    data() {
      return {
        wgsLng: 114.64949701224843,
        wgsLat: 37.851349609806 ,
        gcjLng: 0,
        gcjLat: 0
      };
    },
    methods: {
      convertCoordinates() {
        // 将WGS-84坐标转换为GCJ-02坐标
        // const gcj = coordtransform.wgs84togcj02(this.wgsLng, this.wgsLat);
        const gcj = coordtransform.gcj02tobd09(this.wgsLng, this.wgsLat); 
        this.gcjLng = gcj[0];
        this.gcjLat = gcj[1];
      }
    },
    mounted() {
      // 在组件挂载后执行坐标转换
      this.convertCoordinates();
    }
  };
  </script>
  