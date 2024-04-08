<template>
  <div>
    <el-card id="list">
      <el-row :gutter="15">
        <el-col :span="6">
          <el-input v-model="searchData.title" placeholder="任务名" @change="searchTask()" clearable="">
            <el-button slot="append" icon="el-icon-search"></el-button>
          </el-input>
        </el-col>
        <!-- <el-col :span="6"><el-button type="primary" @click="addTask()">新增</el-button></el-col> -->
      </el-row>


      <!-- 任务列表 -->
      <el-table :data="tableData" style="width: 100%" border="" stripe="">
        <!-- <el-table-column prop="ID" label="ID"></el-table-column> -->
        <el-table-column prop="title" label="任务标题"></el-table-column>
        <el-table-column prop="tag" label="任务标签"></el-table-column>
        <el-table-column prop="briefing" label="任务简介"></el-table-column>
        <el-table-column prop="status" label="任务状态" :formatter="statusFor"></el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button @click="goToDetail(scope.row)" type="text" size="small">详细页面</el-button>
            <el-button v-if="scope.row.status == 0" type="text" size="small" @click="goToEditTask(scope.row)">编辑</el-button>
            <el-button v-if="scope.row.status == 0" @click="deleteTask(scope.row)" type="text" size="small">删除</el-button>
            <el-button v-if="scope.row.status >1" @click="download(scope.row)" type="text" size="small">下载</el-button>
            <el-button v-if="scope.row.status == 2 || scope.row.status == 6" @click="finish(scope.row)" type="text" size="small">完成</el-button>
            <el-button v-if="scope.row.status == 2 || scope.row.status == 6" @click="rejectionBtn(scope.row)" type="text" size="small">退回</el-button>
            <el-button v-if="scope.row.status == 3"  @click="taskCommentBtn(scope.row)" type="text" size="small">评论</el-button>
            <el-button v-if="scope.row.status == 4"  @click="taskCommentBtn(scope.row)" type="text" size="small">修改评论</el-button>

            <!-- 评论对话框 -->
            <el-dialog title="评论" :visible.sync="dialogVisible" width="40%" @close="commentFormLogClose">

              <el-form ref="commentForm" :model="commentForm" label-width="80px">
                <el-form-item label="">
                  <!-- 评分 -->
                  <p>评分</p><el-rate v-model="commentForm.score"></el-rate>
                </el-form-item>

                <!-- 评论内容 -->
                <el-form-item label="">
                  <p>评论内容</p>
                  <el-input type="textarea" v-model="commentForm.content" style="width: 50%;"></el-input>
                </el-form-item>
              </el-form>

              <span slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="addTaskComment()">确 定</el-button>
              </span>
            </el-dialog>

            <!-- 退回对话框 -->
          <el-dialog title="退回任务" :visible.sync="rejectionDialogVisible" width="40%" @close="commentFormLogClose">

              <el-form ref="commentForm" :model="commentForm" label-width="80px">

                <!-- 退回的原因 -->
                <el-form-item label="">
                  <p>退回的原因</p>
                  <el-input type="textarea" v-model="commentForm.content" style="width: 50%;"></el-input>
                </el-form-item>
              </el-form>

              <span slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="rejectionTask()">确 定</el-button>
              </span>
            </el-dialog>

          </template>
        </el-table-column>
      </el-table>
      <!-- 页数 -->
      <div class="pag">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" @size-change="handleSizeChange"
          @current-change="handleCurrentChange" :page-sizes="[1, 2, 3, 6]" :page-size="searchData.pageSize"
          :total="total">
        </el-pagination>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      // 已登录的用户信息
      headers:{
        user_name: '',
        user_id: 0
      },
      // 获取任务列表时的参数
      searchData: {
        // title: "",
        id: 0,
        pageSize: 6,
        pageNum: 1,
      },
      // 获取的任务列表
      tableData: [],
      // 获取的任务总数
      total: 0,
      // 将要添加评论的任务信息
      task:{},
      // 评论对话框状态
      dialogVisible: false,
      // 退回對話框狀態
      rejectionDialogVisible: false,
      score: null,
      // 评论表单
      commentForm: {
        // 评论的任务id
        task_id : 0,
        // 评价者id
        appraiser_id : 0,
        // 评价者name
        appraiser_name: 0,
        // 被评价者id
        receiver_id: 0,
        // 被评价者name
        receiver_name : 0,
        // 评论内容
        content: '',
        // 评论分数
        score: null,
        // 状态码 0：未评论 1：这条评论是买家评论的 2：这条评论是卖家评论的 -1：这条评论是退回的信息
        status : 0,
      },
      // 退回表单
      
    };
  },
  created() {
    this.getHeaders();
    this.getTaskList();
  },
  methods: {
    // 点击退回按钮
    rejectionBtn(item){
      this.task = item;
      this.rejectionDialogVisible = true;
    },
    // 退回task
    rejectionTask(){
      this.commentForm.status = -1;
      this.addTaskCommentPort();
    },
    // 改变评论状态
    async changeTaskStatus(){
      this.dialogVisible = false;
      this.rejectionDialogVisible = false;
      const { data: res } = await this.$http.post(
              "task/EditTask",
              this.task
            );
            if (res.state != 200) {
              this.$message.error(res.message);
            } else {
              this.$message({
                type: "success",
                message: "操作成功!",
              });
            }
            // this.$router.go(0);
    },
    // 点击添加评论
    taskCommentBtn(item){
      console.log(item);
      this.task = item;
      this.dialogVisible = true;
    },
    // 添加评论表单中点击确定按钮
    addTaskComment(){
      this.commentForm.status = 2;
      this.addTaskCommentPort();
    },
    // 向服务器添加评论
    async addTaskCommentPort() {
      var item = this.task;
      console.log(item);
      this.commentForm.task_id = item.ID;
      this.commentForm.appraiser_id = this.headers.user_id;
      this.commentForm.appraiser_name = this.headers.user_name;
      this.commentForm.receiver_id = item.recipient_id;
      this.commentForm.receiver_name = item.recipient_name;
      console.log(this,this.commentForm)
      const { data: res } = await this.$http.post(`taskComment/AddTaskComment`, this.commentForm);
      if (res.state != 200) {
        return this.$message.error(res.message);
      } else {
        if (this.commentForm.status == -1){
            this.task.status = 5;
            this.changeTaskStatus();
        }else{
          this.task.status = 4;
          this.changeTaskStatus();
        }
      }
      this.dialogVisible = false;
    },
    // 评论对话框关闭
    commentFormLogClose() {
      this.$refs.commentForm.resetFields();
    },
    // 获取已登录的用户信息
    getHeaders() {
      this.headers.user_name = window.sessionStorage.getItem("user_name");
      this.headers.user_id = parseInt(window.sessionStorage.getItem("user_id"));
    },
    // 前往编辑任务页面
    goToEditTask(item) {
      if (item.status != 0){
        this.$message.error("任务已被接受，不可编辑！");
        return;
      }
      this.$router.push(`/task/edit/${item.ID}`);
    },
    // 前往任务详情页面
    goToDetail(item){
      // console.log(item)
      this.$router.push(`/task/detail/${item.ID}`);
    },
    // 下载文件
    download(item){
      console.log(item.article)
      window.open(item.article)
    },
    // 完成任务
    finish(item){
      console.log(item)
      this.$confirm.confirm("是否完成此任务", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }).then(async () => {
        // 确定完成任务
        item.status = 3;
        const res = await this.$http.post(`task/EditTask`, item);
        if (res.data.state != 200) {
          return this.$message.error(res.data.message);
        } else {
          this.$message({
            type: "success",
            message: "任务已完成!",
          });
        }
        this.getTaskList();
      }).catch(() => {
        this.$message({
          type: "info",
          message: "已取消",
        });
      });
    },
    // 任务状态格式化
    statusFor(row) {
      if (row.status == 0) {
        return "已发布";
      }
      if (row.status == 1) {
        return "已接受";
      }
      if (row.status == 2) {
        return "已提交";
      }
      if (row.status == 3) {
        return "已完成";
      }
      if (row.status == 4) {
        return "已完成";
      }
      if (row.status == -1) {
        return "已取消";
      }
      if (row.status == 5) {
        return "已被退回";
      }
      if (row.status == 6) {
        return "已重新提交";
      }
    },
    async getTaskList() {
      const { data: res } = await this.$http.get(`task/QueryTaskByDemander/${this.searchData.pageSize}/${this.searchData.pageNum}/${this.headers.user_name}`);
      this.tableData = res.data;
      // console.log(res);
      this.total = res.total;
    },
    // 仅在输入框失去焦点或任务按下回车时触发
    async searchTask() {
      const { data: res } = await this.$http.get(`task/QueryTaskByDemander/${this.searchData.pageSize}/${this.searchData.pageNum}/${this.headers.user_name}`, {
        params: {
          title: this.searchData.title,
        },
      });
      this.tableData = res.data;
      // console.log(res);
      this.total = res.total;
    },
    // 点击删除任务按钮函数
    deleteTask(task) {
      console.log(task)
      if (task.status != 0){
        this.$message.error("任务已被接受，不可删除！");
        return;
      }
      this.$confirm.confirm("此操作将永久删除该文件, 是否继续？", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          // 删除任务
          const res = await this.$http.delete(`task/DeleteTask/${task.ID}`);
          if (res.data.state != 200) {
            return this.$message.error(res.data.message);
          } else {
            this.$message({
              type: "success",
              message: "删除成功!",
            });
          }
          this.getTaskList();
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
    },
    // 点击编辑按钮
    clickEdit(task) {
      console.log(task);
      this.editTaskFormVisible = true;
      this.$nextTick(() => {
        this.editTaskData.title = task.title;
        this.editTaskData.ID = task.ID;
        this.editTaskData.role = task.role;
      });
    },
    handleSizeChange(val) {
      this.searchData.pageSize = val;
      this.getTaskList();
    },
    // table页数改变
    handleCurrentChange(val) {
      this.searchData.pageNum = val;
      this.getTaskList();
    },
  },
};
</script>

<style scoped>
#list {
  padding-bottom: 500px;
}
</style>
