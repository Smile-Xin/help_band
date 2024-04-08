<template>
  <div>
    <el-card>
      <el-form
        ref="task"
        :model="task"
        label-width="80px"
        :rules="taskRules"
      >
        <el-form-item label="任务标题" prop="title">
          <el-input
            v-model="task.title"
            placeholder=""
            style="width: 350px"
          ></el-input>
        </el-form-item>
        
        <el-form-item label="任务类型" prop="tag">
          <el-input
            v-model="task.tag"
            placeholder=""
            style="width: 350px"
          ></el-input>
        </el-form-item>

        <el-form-item label="任务简介">
          <el-input
            type="textarea"
            v-model="task.briefing"
            placeholder=""
            style="width: 350px"
          ></el-input>
        </el-form-item>

        <el-form-item label="任务内容" prop="content">
          <el-input
            type="textarea"
            v-model="task.content"
            placeholder=""
            style="width: 350px"
          ></el-input>
        </el-form-item>
        <div class="but">
          <el-button type="primary" @click="addtask('task')"
            >提交</el-button
          >
        </div>
      </el-form>
    </el-card>

  </div>
</template>

<script>
export default {
  props: ["id"],
  data() {
    return {
      task: {
        tag: "",
        id: 0,
        title: "",
        briefing: "",
        demander_id: 0,
        demander_name: "",
        recipient_id: 0,
        recipient_name: "",
        content: "",
        status: 0,
      },
      taskRules: {
        title: [{ required: true, message: "标题不可为空", trigger: "blur" }],
        tag: [
          { required: true, message: "分类不可为空", trigger: "blur" }
        ],
        content: [
          { required: true, message: "任务内容不可为空", trigger: "blur" }
        ],
      },
    };
  },
  mounted() {
    this.gettask(this.id);
    this.task.demander_id = parseInt(window.sessionStorage.getItem("user_id"));
    this.task.demander_name = window.sessionStorage.getItem("user_name");
    if (this.task.demander_id == null || this.task.demander_id == 0 || this.task.demander_id == "undefined") {
      this.$router.push("/login");
    }
    const token = window.sessionStorage.getItem("token");
    if (token == null || token == "undefined") {
      this.$router.push("/login");
    }
    
  },
  methods: {
    // 获取需要修改的任务
    async gettask(id) {
      if (!id) return;
      const { data: res } = await this.$http.get(`task/QueryTaskByID/${id}`);
      if (res.state != 200) {
        return this.$message.error(res.data.$message);
      }
      this.task = res.data;
      // console.log("this.task", this.task);
    },
    // 添加或者修改任务
    addtask(formName) {
      this.$refs[formName].validate(async (valid) => {
        if (valid) {
          if (this.id) {
            // 修改任务
            const { data: res } = await this.$http.post(
              "task/EditTask",
              this.task
            );
            if (res.state != 200) {
              this.$message.error(res.message);
            } else {
              this.$message({
                type: "success",
                message: "提交成功",
              });
            }
          } else {
            // 添加任务
            console.log("this.id", this.id);
            console.log("this.task", this.task);
            const { data: res } = await this.$http.post(
              "task/AddTask",
              this.task
            );
            if (res.state != 200) {
              this.$message.error(res.message);
            } else {
              this.$message({
                type: "success",
                message: "提交成功",
              });
            }
          }
          this.$router.push("/index");
        } else {
          this.$message({
            message: "格式错误",  
            type: "error",
            duration: 1750,
          });
        }
      });
    },
  },
};
</script>

<style>
.but {
  display: flex;
  justify-content: right;
  margin: 10px 50px;
}

.el-form-item {
  margin-bottom: 70px;
  margin-top: 40px;
}

</style>
