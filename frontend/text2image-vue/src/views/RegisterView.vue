<template>
  <div class="register-container">
    <el-card class="box-card">
      <h2>注册</h2>
      <el-form
        :model="registerForm"
        status-icon
        :rules="rules"
        ref="registerForm"
        label-position="left"
        label-width="80px"
        class="register-form"
      >
        <el-form-item label="用户名" prop="uname">
          <el-input v-model="registerForm.uname"></el-input>
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="registerForm.email"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="pass">
          <el-input
            type="password"
            v-model="registerForm.pass"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="确认密码" prop="password">
          <el-input
            type="password"
            v-model="registerForm.password"
            autocomplete="off"
          ></el-input>
        </el-form-item>
      </el-form>
      <div class="btn-group">
        <el-button type="primary" @click="submitForm('registerForm')">提交</el-button>
        <el-button @click="resetForm('registerForm')">重置</el-button>
        <el-button @click="goBack">返回</el-button>
      </div>
    </el-card>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    const validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'));
      } else if (value.length < 6 || value.length > 16) {
        callback(new Error('密码长度必须在6到16个字符之间'));
      } else {
        if (this.registerForm.password !== '') {
          this.$refs.registerForm.validateField('password');
        }
        callback();
      }
    };
    var validatePass2 = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请再次输入密码"));
      } else if (value !== this.registerForm.pass) {
        callback(new Error("两次输入密码不一致!"));
      } else {
        callback();
      }
    };
    return {
      registerForm: {
        uname: "",
        email: "",
        pass: "",
        password: "",
      },
      rules: {
        uname: [{ required: true, message: "请输入用户名", trigger: "blur" }],
        email: [
          { required: true, message: "请输入邮箱地址", trigger: "blur" },
          { type: "email", message: "请输入正确的邮箱地址", trigger: ["blur", "change"] },
        ],
        pass: [{ required: true, validator: validatePass, trigger: "blur" }],
        password: [
          { required: true, validator: validatePass2, trigger: "blur" },
        ],
      },
    };
  },
  methods: {
    async hashPassword(password) {
      const encoder = new TextEncoder();
      const data = encoder.encode(password);
      const hashBuffer = await crypto.subtle.digest('SHA-256', data);
      const hashArray = Array.from(new Uint8Array(hashBuffer));
      const hashHex = hashArray.map(byte => byte.toString(16).padStart(2, '0')).join('');
      return hashHex;
    },
    async submitForm(formName) {
      this.$refs[formName].validate(async (valid) => {
        if (valid) {
          try {
            const hashedPassword = await this.hashPassword(this.registerForm.password);
            const formattedFormData = {
              email: this.registerForm.email,
              username: this.registerForm.uname,
              password: hashedPassword
            };
            console.log(formattedFormData);
            const response = await axios.post('http://localhost:8080/register', formattedFormData);
            if (response.data.message === '注册成功') {
              console.log('注册成功');
              this.$router.push('/login');
            } else {
              this.$message.error(response.data.message);
            }
          } catch (error) {
            console.error('注册失败:', error);
            if (error.response) {
              this.$message.error(`服务器错误: ${error.response.status} - ${error.response.data.message}`);
            } else if (error.request) {
              this.$message.error('请求未响应，请检查网络连接');
            } else {
              this.$message.error('请求发送失败，请稍后再试');
            }
          }
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    goBack() {
      if (this.$route.path !== '/login') {
        this.$router.push('/login');
      } else {
        this.$router.go(-1);
      }
    }
  }
};
</script>

<style scoped>
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f0f2f5; /* 纯色背景 */
  background-image: url(https://chuhsing-blog-bucket.oss-cn-shenzhen.aliyuncs.com/chuhsing/202411251503682.png);
  background-size: cover;
  background-position: center;
}

.box-card {
  width: 400px;
  padding: 20px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  background-color: #ffffff;
}

h2 {
  text-align: center;
  margin-bottom: 20px;
  color: #333333;
}

.register-form {
  display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    width: 100%;
    height: 100%;
}

.btn-group {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
}

.el-button {
  flex: 1;
  margin: 0 5px;
}

.el-button + .el-button {
  margin-left: 10px;
}
.el-input {
  width: 100%;
  margin-bottom: 10px;
  border-radius: 4px;
  transition: border-color 0.3s ease;
}
</style>