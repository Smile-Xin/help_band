<template>
  <div>
    <v-app-bar mobileBreakpoint="sm" app dark flat color="black darken-2">
      <v-app-bar-nav-icon dark class="hidden-md-and-up" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title v-if="headers.user_name">
        <v-app-bar-nav-icon class="mx-15 hidden-md-and-down">
          <v-avatar size="40" color="grey">
            <img  :src="avatar"  alt />
          </v-avatar>
        </v-app-bar-nav-icon>
      </v-toolbar-title>

      <v-tabs dark center-active centered class="hidden-sm-and-down" key="$route.fullPath" v-model="selectedTab">
        <v-tab to="/index">index</v-tab>
        <v-tab v-if="headers.user_name" to="/task/add">发布任务</v-tab>
        <v-tab v-if="headers.user_name" to="/task/publishedTask">发布的任务</v-tab>
        <v-tab v-if="headers.user_name" to="/task/acceptedTask">接受的任务</v-tab>
        <v-tab v-if="headers.user_name" to="/community">社区</v-tab>
        <v-tab v-if="headers.user_name" :to="`/letter`">私信</v-tab>

        <!-- <v-tab v-if="headers.user_name" to="/task/finishedTask">完成的任务</v-tab> -->
        <!-- <v-tab v-if="headers.user_name" to="/task/myTask">我的任务</v-tab>
        <v-tab v-if="headers.user_name" to="/user/myInfo">个人中心</v-tab>
        <v-tab v-if="headers.user_name" to="/user/editInfo">修改信息</v-tab>
        <v-tab v-if="headers.user_name" to="/user/changePassword">修改密码</v-tab>
        <v-tab v-if="headers.user_name" to="/user/message">消息中心</v-tab>
        <v-tab v-if="headers.user_name" to="/user/help">帮助中心</v-tab> -->
        <!-- <v-tab v-for="item in cateList" :key="item.ID" :to="`/category/${item.ID}`" text>{{ item.name }}</v-tab> -->
      </v-tabs>

      <v-spacer></v-spacer>

      <!-- <v-responsive class="hidden-sm-and-down" color="white">
        <v-text-field dense flat hide-details solo-inverted rounded placeholder="请输入文章标题查找" dark
          append-icon="mdi-text-search" v-model="searchName" @click:append="searchTitle(searchName)"></v-text-field>
      </v-responsive> -->

      <!-- 判断是否登陆了 -->
      <v-btn v-if="!headers.user_name" text dark @click="goLogin()">请登录</v-btn>
      <v-btn v-if="headers.user_name" text dark @click="goToAdmin()">{{ headers.user_name }}</v-btn>
      <v-btn class="hidden-md-and-down" v-if="headers.user_name" text dark @click="loginout">退出</v-btn>
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" color="indigo" dark app temporary>
      <v-list>
        <v-list-item-title>
          <v-btn to="/index" dark text>
            <v-icon small>mdi-home</v-icon>首页
          </v-btn>
        </v-list-item-title>

        <v-list-item v-model="group" active-class="deep-purple--text text--accent-4" v-for="item in cateList"
          :key="item.ID">
          <v-list-item-title>
            <v-btn dark text @click="gotoCate(item.ID)">{{ item.name }}</v-btn>
          </v-list-item-title>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </div>
</template>

<script>
export default {
  data() {
    return {
      // 头像
      avatar : "",
      // 选择的tab
      selectedTab: null,
      drawer: false,
      group: null,
      valid: true,
      registerformvalid: true,
      cateList: [],
      searchName: '',
      formdata: {
        user_name: '',
        password: ''
      },
      checkPassword: '',
      dialog: false,
      headers: {
        Authorization: '',
        user_name: ''
      },
      nameRules: [
        (v) => !!v || '用户名不能为空',
        (v) =>
          (v && v.length >= 4 && v.length <= 12) ||
          '用户名必须在4到12个字符之间'
      ],
      passwordRules: [
        (v) => !!v || '密码不能为空',
        (v) =>
          (v && v.length >= 6 && v.length <= 20) || '密码必须在6到20个字符之间'
      ],
      checkPasswordRules: [
        (v) => !!v || '密码不能为空',
        (v) =>
          (v && v.length >= 6 && v.length <= 20) || '密码必须在6到20个字符之间',
        (v) => v === this.formdata.password || '密码两次输入不一致，请检查'
      ]
    }
  },
  watch: {
    group() {
      this.drawer = false
    }
  },
  created() {
    // this.cteatedFun()
  },
  mounted() {
    // console.log(window.sessionStorage.getItem('token'))
    // console.log(window.sessionStorage.getItem('user_name'))
    this.headers = {
      Authorization: `Bearer ${window.sessionStorage.getItem('token')}`,
      user_name: window.sessionStorage.getItem('user_name'),
      user_id : parseInt(window.sessionStorage.getItem("user_id"))

    }
    this.getProfileInfo()
  },
  methods: {
    // 获取头像
    async  getProfileInfo() {
      if (!this.headers.user_name) return
       const { data: res } = await this.$http.get(`admin/profile/${this.headers.user_id}`)
       if (res.state !== 200) return this.$message.error(res.message)
       this.avatar = res.data.avatar
      //  console.log("res",res)
    },
    // 前往后台管理页面
    goToAdmin() {
      let data = this.$router.resolve('/admin')
      window.open(data.href, '_blank')
    },
    // 获取分类
    cteatedFun() {
      console.log("this.$router.path",this.$route.path)
    },

    // 查找文章标题
    searchTitle(title) {
      if (title.length == 0) return this.$message.error('你还没填入搜索内容哦')
      this.selectedTab = 'index'
      this.$router.push(`/search/${title}`)
    },

    gotoCate(cid) {
      // console.log(cid)
      this.$router.push(`/category/${cid}`).catch((err) => err)
    },
    // 跳转到登录页面
    goLogin() {
      window.sessionStorage.setItem('route', this.$route.path)
      this.$router.push('/login')
    },

    // 登录
    async login() {
      if (!this.$refs.loginFormRef.validate())
        return this.$message.error('输入数据非法，请检查输入的用户名和密码')
      const { data: res } = await this.$http.post('loginfront', this.formdata)
      if (res.state !== 200) return this.$message({
        message: res.message,
        type: 'error'
      })
      // console.log(res)
      window.sessionStorage.setItem('user_name', res.name)
      window.sessionStorage.setItem('user_id', res.id)
      this.$message({
        message: '登录成功',
        type: 'success',
        duration: 1750
      })
      this.dialog = false

      this.$router.go(0)
    },

    // 退出
    loginout() {
      window.sessionStorage.clear('token')
      window.sessionStorage.clear('user_name')
      this.$message.success('退出成功')
      this.$router.go(0)
    },

    // 注册
    async registerUser() {
      if (!this.$refs.registerformRef.validate())
        return this.$message.error('输入数据非法，请检查输入的用户名和密码')
      const { data: res } = await this.$http.post('user/add', {
        user_name: this.formdata.user_name,
        password: this.formdata.password,
        role: 1
      })
      if (res.state !== 200) return this.$message.error(res.message)
      this.$message.success('注册成功')
      this.dialog = false
      this.$router.go(0)
    }
  }
}
</script>

<style></style>
