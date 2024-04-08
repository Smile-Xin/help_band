<template>
  <div class="container">
    <div class="loginBox">

      <el-form :model="loginForm" ref="loginForm" :rules="loginRules" label-width="80px" class="loginForm">
        <el-button id="back" type="danger" @click="back()" circle><i class="el-icon-close"></i></el-button>
        <el-form-item label="用户名" size="normal" prop="user_name">
          <el-input v-on:keyup.enter.native="login('loginForm')" v-model="loginForm.user_name" placeholder="user_name"
            size="normal" prefix-icon="el-icon-user"></el-input>
        </el-form-item>

        <el-form-item label="密码" size="normal" prop="password">
          <el-input v-on:keyup.enter.native="login('loginForm')" v-model="loginForm.password" type="password"
            placeholder="password" size="normal" prefix-icon="el-icon-lock"></el-input>
        </el-form-item>

        <el-form-item label="" size="normal">
          <el-button type="primary" @click="signUp('loginForm')">注册</el-button>
          <el-button type="" @click="resetForm('loginForm')">清除</el-button>
        </el-form-item>

      </el-form>

    </div>
  </div>
</template>

<script>

export default {
  data() {
    return {
      url: "",
      loginForm: {
        user_name: '',
        password: ''
      },
      loginRules: {
        user_name: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 3, max: 10, message: '长度在 3 到 10 个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 3, max: 15, message: '长度在 3 到 15 个字符', trigger: 'blur' }
        ]
      }
    };
  },
  mounted() {
    // 获取上一个页面的路由，在登陆成功后跳转回去。由TopBar.vue传递过来
    this.url = window.sessionStorage.getItem("route")
    if (this.url == 'undefined' || this.url == null || this.url == "") {
      this.url = "/index"
    }
  },
  methods: {
    // 重置表单
    resetForm(formName) {
      //  this.$refs.loginForm.resetFields();
      this.$refs[formName].resetFields();
    },
    // 注册
    signUp(formName) {
      this.$refs[formName].validate(async valid => {
        if (valid) {
          console.log("this.loginForm", this.loginForm);
          // 调用注册接口
          const { data: res } = await this.$http.post("user/AddUser", this.loginForm)
          // 判断是否注册成功
          if (res.state != 200) { return this.$message.error(res.message) }
          // 注册成功
          this.$message({
            type: 'success',
            message: "注册成功"
          })
          this.$router.push(this.url)
        } else {
          this.$message({
            message: "格式错误",
            type: 'error',
            duration: 1750
          })
        }
      })
    },
    // 返回
    back() {
      this.$router.push(this.url)
    }
  },


}
</script>

<style scoped>
#back {
  position: absolute;
  top: -80px;
  right: 0;
  background-color: #fff;
  margin: 0;
  border: 0;

}

.el-icon-close {
  color: black;
  font-size: 20px;
}

.container {
  height: 100%;
  max-width: 100%;
  background-color: #c2c2c2;

  /* flex布局兼容性较差，这里仅作参考 */
  /* display: flex;
  justify-content: center;
  align-items: center; */

}

.loginBox {
  width: 450px;
  height: 300px;
  background-color: #fff;

  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 5px;
}

.loginForm {
  width: 100%;
  position: absolute;
  bottom: 10%;
  padding: 0 20px;
  box-sizing: border-box;

}

html {
  overflow-y: hidden;
}
</style>