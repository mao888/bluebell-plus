<template>
  <div class="go-to-top">
    <button class="top-btn">back to top</button>
  </div>
</template>>


<script>
export default {
  name: "ToTop",
};
</script>

<style lang="less" scoped>
.top-btn{
        width: 128px;
        height: 32px;
        position: fixed;
        bottom: 10px;
        right: 20px;
        background-color: #0079d3;
        color: #fff;
        border-radius: 4px;
        font-size: 14px;
    }
</style>