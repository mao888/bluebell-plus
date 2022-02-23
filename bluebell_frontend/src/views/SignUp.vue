<template>
  <div class="main">
    <div class="container">
      <h2 class="form-title">注册</h2>
      <div class="form-group">
        <label for="name">用户名</label>
        <input type="text" class="form-control" name="name" id="name" placeholder="用户名" v-model="username"/>
      </div>
      <div class="form-group">
        <label for="pass">密码</label>
        <input type="password" class="form-control" name="pass" id="pass" placeholder="密码" v-model="password"/>
      </div>
      <div class="form-group">
        <label for="re_pass">确认密码</label>
        <input type="password" class="form-control" name="re_pass" id="re_pass" placeholder="确认密码"  v-model="re_password"/>
      </div>
      <div class="form-btn">
        <button type="button" class="btn btn-info" @click="submit">提交</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
	name: "SignUp",
	data() {
		return {
			username: "",
			password: "",
			re_password: "",
			submitted: false
		};
	},
	computed: {
	},
	created() {

	},
	methods: {
		submit() {
			this.$axios({
				method: 'post',
				url:'/signup',
				data: JSON.stringify({
					username: this.username,
					password: this.password,
					re_password: this.re_password
				})
			}).then((res)=>{
				console.log(res.data);
				if (res.code == 1000) {
          console.log('signup success');
          this.$router.push({ name: "Login" });
				}else{
          console.log(res.msg);
        }
			}).catch((error)=>{
				console.log(error)
			})
		}
	}
};
</script>
<style lang="less" scoped>
.main {
  background: #f8f8f8;
  padding: 150px 0;
  .container {
    width: 600px;
    background: #fff;
    margin: 0 auto;
    max-width: 1200px;
    padding: 20px;
    .form-title {
      margin-bottom: 33px;
      text-align: center;
    }
    .form-group {
      margin: 15px;
      label {
        display: inline-block;
        max-width: 100%;
        margin-bottom: 5px;
        font-weight: 700;
      }
      .form-control {
        display: block;
        width: 100%;
        height: 34px;
        padding: 6px 12px;
        font-size: 14px;
        line-height: 1.42857143;
        color: #555;
        background-color: #fff;
        background-image: none;
        border: 1px solid #ccc;
        border-radius: 4px;
      }
    }
    .form-btn {
      display: flex;
      justify-content: center;
      .btn {
        padding: 6px 20px;
        font-size: 18px;
        line-height: 1.3333333;
        border-radius: 6px;
        display: inline-block;
        margin-bottom: 0;
        font-weight: 400;
        text-align: center;
        white-space: nowrap;
        vertical-align: middle;
        -ms-touch-action: manipulation;
        touch-action: manipulation;
        cursor: pointer;
        border: 1px solid transparent;
      }
      .btn-info {
        color: #fff;
        background-color: #5bc0de;
        border-color: #46b8da;
      }
    }
  }
}
</style>