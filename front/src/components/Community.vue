<template>
  <v-col class="blue lighten-5 list">
    <v-card class="ma-3 d-flex flex-no-wrap justify-space-between align-center" v-for="item in artList" :key="item.id"
      link @click="$router.push(`article/detail/${item.ID}`)">
      <!-- <v-avatar class="ma-3 hidden-sm-and-down" size="125" tile>
          <v-img :src="item.img"></v-img>
        </v-avatar> -->

      <v-col class="">
        <v-card-title>
          <v-chip color="pink" outlined label class="mr-3 white--text">
            {{
              item.Category.name
            }}
          </v-chip>
          <div>{{ item.title }}</div>
        </v-card-title>
        <v-card-subtitle class="mt-1" v-text="item.desc"></v-card-subtitle>
        <v-divider class="mx-4"></v-divider>
        <v-card-text class="d-flex align-center">
          <div class="d-flex align-center">
            <v-icon class="mr-1" small>{{ 'mdi-calendar-month' }}</v-icon>
            <span>{{ item.UpdatedAt | dateformat('YYYY-MM-DD') }}</span>
          </div>
          <div class="mx-4 d-flex align-center">
            <v-icon class="mr-1" small>{{ 'mdi-comment' }}</v-icon>
            <span>{{ item.comment_count }}</span>
          </div>
          <div class="mx-1 d-flex align-center">
            <v-icon class="mr-1" small>{{ 'mdi-eye' }}</v-icon>
            <span>{{ item.read_count }}</span>
          </div>
          <div @click.stop="openProfileDialog(item.demander_name)" class="mx-1 d-flex align-center">
            <span>作者：{{ item.User.user_name }}</span>
          </div>
        </v-card-text>
      </v-col>
    </v-card>
    <div class="text-center">
      <v-pagination color="black" total-visible="7" v-model="queryParam.pagenum"
        :length="Math.ceil(total / queryParam.pagesize)" @input="getArtList()"></v-pagination>
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
                    {{ score }}
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
            </v-card>
          </el-dialog>
  </v-col>
</template>
<script>
export default {
  data() {
    return {
      // 个人信息
      profile: {},
      score: 0,
      // 个人信息弹出框
      profileDialog: false,
      artList: [],
      queryParam: {
        pagesize: 5,
        pagenum: 1
      },

      count: 0,
      total: 0
    }
  },
  created() { },
  mounted() {
    this.getArtList()
  },
  filters: {},
  methods: {
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
      this.profile = res.data.profile
      this.score = res.data.score
      console.log("this.profile", this.profile)
      console.log("score", this.score)
    },
    // 获取文章列表
    async getArtList() {
      const { data: res } = await this.$http.get('article/get', {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      this.artList = res.data
      this.total = res.total
      console.log("this.artList", this.artList)
    }
  }
}
</script>
<style scoped>
.list {
  padding-bottom: 500px;
}
</style>