<template>
  <el-container>

    <el-container>
      <el-aside width="300px" style="height: 600px;">
        <el-row class="tac">
          <el-menu default-active="" router="" mode="vertical" class="el-menu-vertical-demo" background-color="#d3dce6"
            text-color="#282c34" active-text-color="#004d40" style="padding: 0%;">

            <el-menu-item v-for="item in letterList" :key="item.id" :route="`chat/${item.id}`" style="padding: 0%;">
              <template slot="title">
                <i class="el-icon-pie-chart"></i>
                <span>{{item}}</span>
              </template>
            </el-menu-item>



          </el-menu>
        </el-row>
      </el-aside>
      <el-main>
        <router-view :key="$route.path"></router-view>
      </el-main>
    </el-container>

  </el-container>
</template>

<script>
export default {
  props: ["userId"],
  data() {
    return {
      // letterList
      letterList: [],
      //用户信息
      headers:{},
      def: "2-1",
    };
  },
  created() {
    this.getLetterList();
    this.headers.user_id = window.sessionStorage.getItem("user_id");
    this.headers.user_name = window.sessionStorage.getItem("user_name");
  },
  methods: {
    // 获取信件列表
    async getLetterList() {
      const { data: res } = await this.$http.get(`letter/QueryLetterByUserId/${this.headers.user_id}`)
      console.log("res", res)
      if (res.state !== 200) {
        this.$message.error(res.message);
        return;
      }
      this.letterList = res.data;
    },
  },
}
</script>

<style scope></style>
