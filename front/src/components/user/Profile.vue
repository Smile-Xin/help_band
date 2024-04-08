<template>
  <el-card>
    <el-form
      labelAlign="left"
      :label-col="{ span: 2 }"
      :wrapper-col="{ span: 12 }"
    >
      <!-- <el-form-item label="昵称">
        <el-input style="width: 300px" v-model="profileInfo.name"></el-input>
      </el-form-item> -->

      <el-form-item label="个人简介">
        <el-input type="textarea" v-model="profileInfo.desc"></el-input>
      </el-form-item>

      <el-form-item label="QQ号码">
        <el-input style="width: 300px" v-model="profileInfo.qq_chat"></el-input>
      </el-form-item>

      <el-form-item label="微信">
        <el-input style="width: 300px" v-model="profileInfo.wechat"></el-input>
      </el-form-item>

      <el-form-item label="微博">
        <el-input style="width: 300px" v-model="profileInfo.weibo"></el-input>
      </el-form-item>

      <el-form-item label="B站地址">
        <el-input style="width: 300px" v-model="profileInfo.bili"></el-input>
      </el-form-item>

      <el-form-item label="Email">
        <el-input style="width: 300px" v-model="profileInfo.email"></el-input>
      </el-form-item>

      <el-form-item label="头像">
        <el-upload style="display: inline-block;" class="upload-demo" :action="uploadUrl" :on-success="handleSuccess" :on-error="handleError"
              :show-file-list=false :headers="headerObj">
              <el-button size="small" type="primary" @click="upBtn">点击上传</el-button>
            </el-upload>
      </el-form-item>

      <el-form-item>
        <el-button
          type="danger"
          style="margin-right: 15px"
          @click="updateProfile"
          >更新</el-button
        >
      </el-form-item>
    </el-form>
  </el-card>
</template>
<script>
// import { Url } from '../../plugin/http'

export default {
  data() {
    return {
      profileInfo: {
        id: 0,
        name: "",
        desc: "",
        qq_chat: "",
        wechat: "",
        weibo: "",
        bili: "",
        email: "",
        img: "",
        avatar: "",
        icp_record: "",
      },
      uploadUrl: "http://localhost:3001/api/v1/qiniu/UploadAvatar/",
      headers: {},
      headerObj: {
        "Authorization": "Bearer " + window.sessionStorage.getItem("token"),
      },
    };
  },
  created() {
    
    this.headers.user_name = window.sessionStorage.getItem("user_name");
    this.headers.user_id = parseInt(window.sessionStorage.getItem("user_id"));
    this.getProfileInfo();
  },
  methods: {
    // 上传按钮
    upBtn() {
      this.uploadUrl = `http://localhost:3001/api/v1/qiniu/uploadAvatar/${this.headers.user_name}`;
    },
    // 上传文件成功的钩子
    handleSuccess(res, file) {
      console.log(res);
      this.profileInfo.avatar = res.url;
      console.log(file);
      this.$message({
        message: "上传成功",
        type: "success",
      });
      // this.$router.go(0);
    },
    // 上传文件失败的钩子
    handleError(err, file) {
      console.log(err);
      console.log(file);
      this.$message.error("上传失败");
    },
    // 获取个人设置
    async getProfileInfo() {
      const { data: res } = await this.$http.get(
        `admin/profile/${this.headers.user_id}` 
      );
      console.log(res);
      if (res.state !== 200) {
        // if (res.status === 1004 || 1005 || 1006 || 1007) {
        //   window.sessionStorage.clear()
        //   this.$router.push('/login')
        // }
        this.$message.error(res.message);
      }
      this.profileInfo = res.data;
    },
    // 上传头像
    avatarChange(info) {
      console.log(this.upUrl);
      if (info.file.status !== "uploading") {
        console.log(info.file, info.fileList);
      }
      if (info.file.status === "done") {
        this.$message.success(`图片上传成功`);
        const imgUrl = info.file.response.url;
        this.profileInfo.avatar = imgUrl;
      } else if (info.file.status === "error") {
        this.$message.error(`图片上传失败`);
      }
    },
    // // 上传头像背景图
    // imgChange(info) {
    //   if (info.file.status !== 'uploading') {
    //   }
    //   if (info.file.status === 'done') {
    //     this.$message.success(`图片上传成功`)
    //     const imgUrl = info.file.response.url
    //     this.profileInfo.img = imgUrl
    //   } else if (info.file.status === 'error') {
    //     this.$message.error(`图片上传失败`)
    //   }
    // },

    // 更新
    async updateProfile() {
      const { data: res } = await this.$http.post(
        `profile/edit/${this.profileInfo.id}`,
        this.profileInfo
      );
      if (res.state !== 200) return this.$message.error(res.message);
      this.$message.success(`个人信息更新成功`);
      this.$router.push("/admin");
    },
  },
};
</script>

<style scoped>
.upBtn {
  margin-right: 10px;
}
</style>
