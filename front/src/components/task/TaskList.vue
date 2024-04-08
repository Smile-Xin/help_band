<template>
  <v-col class="blue lighten-5 list">
    <v-card class="ma-3 d-flex flex-no-wrap justify-space-between align-center" v-for="item in taskList" :key="item.id"
      link @click="$router.push(`task/detail/${item.ID}`)">
      <!-- <v-avatar class="ma-3 hidden-sm-and-down" size="125" tile>
        <v-img :src="item.img"></v-img>
      </v-avatar> -->

      <v-col class="">
        <v-card-title>
          <!-- 表示类型的小粉框框 -->
          <!-- <v-chip color="pink" outlined label class="mr-3 white--text">
            {{
            item.tag
            }}
          </v-chip> -->

          <div>{{ item.title }}</div>
        </v-card-title>
        <v-card-subtitle class="mt-1" v-text="item.tag"></v-card-subtitle>
        <v-card-subtitle class="mt-1" v-text="item.briefing"></v-card-subtitle>
        <!-- 时间上的那条横线 -->
        <v-divider class="mx-4"></v-divider>

        <v-card-text class="d-flex align-center">
          <div class="d-flex align-center">
            <v-icon class="mr-1" small>{{ 'mdi-calendar-month' }}</v-icon>
            <span>{{ item.CreatedAt | dateformat('YYYY-MM-DD') }}</span>
          </div>

          <!-- 需求方的名字 -->
          <div @click.stop="openProfileDialog(item.demander_name)" class="mx-4 d-flex align-center">
            <v-icon class="mr-1" small>{{ 'mdi-account-question' }}</v-icon>
            <span>{{ item.demander_name }}</span>
          </div>
        </v-card-text>

      </v-col>
    </v-card>
    <br>
    <br>
    <div class="text-center">
      <v-pagination color="black" total-visible="7" v-model="queryParam.pagenum"
        :length="Math.ceil(total / queryParam.pagesize)" @input="getTaskList()"></v-pagination>
    </div>

    <!-- 个人信息弹出框 -->
    <el-dialog class="profile-dialog" :visible.sync="profileDialog" width="500px" :before-close="handleClose">
      <v-card class="mx-auto my-12 profile-card" max-width="100%" max-height="100%">
        <template slot="progress">
          <v-progress-linear color="deep-purple" height="10" indeterminate></v-progress-linear>
        </template>
        <div height="250px" style="text-align: center;background-color: #f5f5f5;">
          <v-avatar size="300" color="grey">
            <img :src="profile.avatar" alt />
          </v-avatar>
        </div>
        <v-card-title>{{ profile.name }}</v-card-title>
        <!-- 评分 -->
        <v-card-text>
          <v-row class="mx-0">
            <v-rating :value="profile.score" color="amber" dense half-increments readonly size="14"></v-rating>
            <div class="grey--text ms-4">
              {{ profile.score }}
            </div>
          </v-row>
          <!-- 副标题 -->
          <div class="my-4 text-subtitle-1">
            <!-- 副标题 -->
          </div>
          <div>{{ profile.desc }}</div>
        </v-card-text>
        <!-- 横线 -->
        <v-divider class="mx-4"></v-divider>
        <v-card-title>联系方式</v-card-title>
        <v-card-text>
          <div class="my-4 text-subtitle-1">微博：{{ profile.weibo }}</div>
          <div class="my-4 text-subtitle-1">QQ：{{ profile.qq_chat }}</div>
          <div class="my-4 text-subtitle-1">Wechat：{{ profile.wechat }}</div>
          <div class="my-4 text-subtitle-1">BiliBili：{{ profile.bili }}</div>
          <div class="my-4 text-subtitle-1">Email：{{ profile.email }}</div>
        </v-card-text>
        <el-button v-if="this.headers.user_id !== profile.ID" class="message-btn" color="primary" text @click="goToMessage()">私信</el-button>
      </v-card>
    </el-dialog>

  </v-col>
</template>


<script>
export default {
  data() {
    return {
      // 个人信息表的数据
      // 个人信息
      profile: {},
      // 个人信息弹出框
      profileDialog: false,
      // user信息
      user: {},
      taskList: [],
      queryParam: {
        pagesize: 5,
        pagenum: 1,
        title: ' '
      },
      headers: {
        user_id: window.sessionStorage.getItem('user_id'),
        user_name: window.sessionStorage.getItem('user_name'),
      },

      count: 0,
      total: 0
    }
  },
  created() { },
  mounted() {
    this.getTaskList()
  },
  filters: {},
  methods: {
    // 跳转到私信页面
    goToMessage() {
      if (this.headers.user_id === 0 || this.headers.user_id === null || this.headers.user_id === "undefined") return alert("请先登录")
      this.$router.push(`/letter/chat/${this.user.ID}`)
    },
    // 关闭个人信息弹出框
    handleClose() {
      this.profileDialog = false
    },
    // 打开个人信息弹出框
    openProfileDialog(demander_name) {
      this.getProfile(demander_name)
      this.profileDialog = true

    },
    // 获取个人信息
    async getProfile(demander_name) {
      const { data: res } = await this.$http.get(`user/QueryByName/${demander_name}`, {})
      console.log("res", res)
      this.profile = res.data.Profile
      this.user = res.data
      this.profile.score = parseFloat(res.data.score).toFixed(1)
      console.log("this.profile", this.profile)
      console.log("this.user", this.user)
    },
    // 获取文章列表
    async getTaskList() {
      // console.log('this.queryParam', this.queryParam)
      const { data: res } = await this.$http.get(`task/GetAll/${this.queryParam.pagesize}/${this.queryParam.pagenum}`, {})
      this.taskList = res.data
      this.total = res.total
      // console.log("total",this.total)
      // console.log("pagesize",this.queryParam.pagesize)
      // console.log("this.taskList",Math.ceil(this.total / this.queryParam.pagesize))
      // console.log("this.taskList",this.taskList)
    }
  }
}
</script>
<style>
.list {
  padding-bottom: 50px;
}

.profile-card {
  max-width: 100%;
  max-height: 100%;
  width: 100%;
  height: 100%;
}

.profile-dialog {
  border: 2px solid;
}

.el-dialog .el-dialog__body {
  padding: 0 0 0 0;
}

.el-dialog__header {

  display: none;
}

.message-btn {
  position: relative;
  left: 215px;
  top: -30px;
  color: black;
}
</style>
