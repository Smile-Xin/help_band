<template>
  <div>
    <el-card id="list">
      <el-row :gutter="15">
        <el-col :span="6">
          <el-input v-model="searchData.title" placeholder="任务名" @change="searchTask()" clearable="">
            <el-button slot="append" icon="el-icon-search"></el-button>
          </el-input>
        </el-col>
      </el-row>


      <!-- 任务列表 -->
      <el-table :data="tableData" style="width: 100%" border="" stripe="">
        <el-table-column prop="title" label="任务标题"></el-table-column>
        <el-table-column prop="tag" label="任务标签"></el-table-column>
        <el-table-column prop="briefing" label="任务简介"></el-table-column>
        <el-table-column prop="demander_name" label="提问者"></el-table-column>
        <el-table-column prop="status" label="任务状态" :formatter="statusFor"></el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button @click="goToDetail(scope.row)" type="text" size="small">详细页面</el-button>
            <el-button @click="dialogVisible = true" type="text" size="small">评论</el-button>

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
                <el-button type="primary" @click="addTaskComment(scope.row)">确 定</el-button>
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
      headers: {
        user_name: '',
        user_id: 0
      },
      // 获取任务列表时的参数
      searchData: {
        title: "",
        id: 0,
        pageSize: 6,
        pageNum: 1,
      },
      // 获取的任务列表
      tableData: [],
      // 获取的任务总数
      total: 0,
      // 评论对话框状态
      dialogVisible: false,
      score: null,
      // 评论表单
      commentForm: {
        // 评论内容
        content: '',
        // 评论分数
        score: null,
        // 任务id
        task_id: 0,
        // 评论者id
        appraiser_id: 0,
        // 
      },
    };
  },
  created() {
    this.getHeaders();
    this.getTaskList();
  },
  methods: {
    // 获取已登录的用户信息
    getHeaders() {
      this.headers.user_name = window.sessionStorage.getItem("user_name");
      this.headers.user_id = parseInt(window.sessionStorage.getItem("user_id"));
    },
    // 前往任务详情页面
    goToDetail(item) {
      // console.log(item)
      this.$router.push(`/task/detail/${item.ID}`);
    },
    // 评论对话框关闭
    commentFormLogClose() {
      this.$refs.commentForm.resetFields();
    },
    // 获取任务列表
    async getTaskList() {
      const { data: res } = await this.$http.get(`task/QueryTaskByUserIdStatus/${this.searchData.pageSize}/${this.searchData.pageNum}/${this.headers.user_name}/3`);
      this.tableData = res.data;
      console.log(res);
      this.total = res.total;
    },
    // 仅在输入框失去焦点或任务按下回车时触发
    async searchTask() {
      console.log(this.searchData.title);
      const { data: res } = await this.$http.get(`task/QueryTaskByUserIdStatus/${this.searchData.pageSize}/${this.searchData.pageNum}/${this.headers.user_name}/3`, {
        params: {
          title: this.searchData.title,
        },
      });
      this.tableData = res.data;
      console.log(res);
      this.total = res.total;
    },
    // 添加任务评论
    // async addTaskComment(item) {
    //   console.log(item);
    //   this.commentForm.score = this.commentForm.score;
    //   this.commentForm.content = this.commentForm.content;
    //   this.commentForm.task_id = item.ID;
    //   this.commentForm.user_id = this.headers.user_id;
    //   this.commentForm.user_name = this.headers.user_name;
    //   console.log(this.commentForm);
    //   const { data: res } = await this.$http.post(`task/AddTaskComment`, this.commentForm);
    //   if (res.state != 200) {
    //     return this.$message.error(res.message);
    //   } else {
    //     this.$message({
    //       type: "success",
    //       message: "评论成功!",
    //     });
    //   }
    //   this.dialogVisible = false;
    // },
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
    },
    // table每页大小改变
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
