<template>
    <div>
        <div ref="myDom">
            <MyHeader />
        </div>
        <!-- <el-scrollbar max-height="100vh" @scroll="headerScroll"> -->
        <div class="bg-#F0F2F5 dark_bg_black pt100">
            <div class="max-w-1280 mx-auto  rounded-t-23 bg-#fff p30 min-h-100vh lt-lg:px15">
                <slot></slot>
            </div>
        </div>
        <MyFooter />
        <!-- </el-scrollbar> -->
        <el-backtop :right="100" :bottom="100" class="lt-lg:important:right-20 lt-lg:important:bottom-50" />
        <!-- <SearchDialog /> -->
    </div>
</template>
<script lang="ts" setup>
const myDom = ref()
let scrollCount = ref(0)
const headerScroll = (res: any) => {
    // const scrollNow = ref(document.documentElement.scrollTop)
    // console.log(Math.abs(scrollNow.value - scrollCount.value));
    if (res.scrollTop > scrollCount.value) {
        myDom.value.firstChild.setAttribute('style', `top:-50px`);
    }
    else {
        myDom.value.firstChild.setAttribute('style', 'top:0px');
    }
    scrollCount.value = res.scrollTop
}
</script>

