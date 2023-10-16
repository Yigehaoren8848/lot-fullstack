<template>
    <div>
        <div style="display: flex;align-items: center;width: 100%;height: 40px;margin-bottom: 0px;">
            <div>
                <img style="width:30px;height: 30px;" src="../../images/go.png" alt="">
            </div>
            <div >
                {{ title }}
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                
            }
        },
        props:{
            title: {
                type:String,
                default:""
            }
        }
    }
</script >

<style lang="scss">

</style>