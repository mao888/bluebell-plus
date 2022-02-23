<template>
  <div id="app">
    <div class="page">
      <HeadBar />
      <router-view />
    </div>
  </div>
</template>


<script>
import HeadBar from "@/components/HeadBar.vue";
export default {
  components: {
    HeadBar
  }
};
</script>

<style lang="less">
@import url("./assets/css/iconfont.css");
html, body{
    width: 100%;
    height: 100%;
    font-family: IBMPlexSans, Arial, sans-serif;
    background: #eeeeee;
}
html, body, div, span, applet, object, iframe, h1, h2, h3, h4, h5, h6, p, blockquote, pre, a, abbr, acronym, address, big, button, cite, code, del, dfn, em, img, input, ins, kbd, q, s, samp, small, strike, strong, sub, sup, tt, var, b, u, i, center, dl, dt, dd, ol, ul, li, fieldset, form, label, legend, table, caption, tbody, tfoot, thead, tr, th, td, article, aside, canvas, details, embed, figure, figcaption, footer, header, hgroup, menu, nav, output, ruby, section, summary, time, mark, audio, video {
    margin: 0;
    padding: 0;
    border: 0;
    outline: 0;
  .page {
    width: 100%;
    height: auto;
  }
}

</style>
