<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="8">
                <el-card shadow="hover" class="mgb20" style="height: 252px">
                    <div class="user-info">
                        <img src="../../assets/img/img.jpg" class="user-avator" alt />
                        <div class="user-info-cont">
                            <div class="user-info-name">{{ name }}</div>
                            <div>{{ role }}</div>
                        </div>
                    </div>
                    <div class="user-info-list">
                        上次登录时间：
                        <span>{{ userInfo.lastLoginTime }}</span>
                    </div>
                    <div class="user-info-list">
                        上次登录 I P &nbsp;：
                        <span>{{ userInfo.lastLoginIp }}</span>
                    </div>
                </el-card>
                <el-card shadow="hover" style="height: 252px">
                    <div slot="header" class="clearfix">
                        <span>语言详情</span>
                    </div>
                    Vue <el-progress :percentage="71.3" color="#42b983"></el-progress>JavaScript
                    <el-progress :percentage="24.1" color="#f1e05a"></el-progress>CSS <el-progress :percentage="13.7"></el-progress>HTML
                    <el-progress :percentage="5.9" color="#f56c6c"></el-progress>
                </el-card>
            </el-col>
            <el-col :span="16">
                <el-row :gutter="20" class="mgb20">
                    <el-col :span="8">
                        <el-card shadow="hover" :body-style="{ padding: '0px' }">
                            <div class="grid-content grid-con-1">
                                <i class="el-icon-lx-people grid-con-icon"></i>
                                <div class="grid-cont-right">
                                    <div class="grid-num">1234</div>
                                    <div>用户访问量</div>
                                </div>
                            </div>
                        </el-card>
                    </el-col>
                    <el-col :span="8">
                        <el-card shadow="hover" :body-style="{ padding: '0px' }">
                            <div class="grid-content grid-con-2">
                                <i class="el-icon-lx-notice grid-con-icon"></i>
                                <div class="grid-cont-right">
                                    <div class="grid-num">321</div>
                                    <div>系统消息</div>
                                </div>
                            </div>
                        </el-card>
                    </el-col>
                    <el-col :span="8">
                        <el-card shadow="hover" :body-style="{ padding: '0px' }">
                            <div class="grid-content grid-con-3">
                                <i class="el-icon-lx-goods grid-con-icon"></i>
                                <div class="grid-cont-right">
                                    <div class="grid-num">5000</div>
                                    <div>数量</div>
                                </div>
                            </div>
                        </el-card>
                    </el-col>
                </el-row>
                <el-card shadow="hover" style="height: 403px">
                    <div slot="header" class="clearfix">
                        <span>待办事项</span>
                        <el-button style="float: right; padding: 3px 0" type="text" @click="addTodo">添加</el-button>
                    </div>
                    <el-table :show-header="false" :data="todoList" style="width: 100%">
                        <el-table-column width="40">
                            <template slot-scope="scope">
                                <el-checkbox v-model="scope.row.finished_status"></el-checkbox>
                            </template>
                        </el-table-column>
                        <el-table-column>
                            <template slot-scope="scope">
                                <div v-if="scope.row.showInput">
                                    <el-input
                                        v-model="scope.row.todo"
                                        @focus="focusEvent(scope.row)"
                                        @blur="handleSaveTodo(scope.row)"
                                        v-focus
                                    ></el-input>
                                </div>
                                <div v-else>
                                    <s v-if="scope.row.finished_status">{{ scope.row.todo }}</s>
                                    <p v-else>{{ scope.row.todo }}</p>
                                </div>
                            </template>
                        </el-table-column>
                        <el-table-column width="120">
                            <template slot-scope="scope">
                                <el-button type="text" icon="el-icon-edit" @click="handleEditTodo(scope.row)">编辑</el-button>
                                <el-button type="text" icon="el-icon-delete" class="red" @click="handleDeleteTodo(scope.row)"
                                    >删除</el-button
                                >
                            </template>
                        </el-table-column>
                    </el-table>
                </el-card>
            </el-col>
        </el-row>
    </div>
</template>

<script>
import Schart from 'vue-schart';
import request from '../../utils/request';
export default {
    name: 'dashboard',
    data() {
        return {
            name: localStorage.getItem('ms_username'),
            userInfo: {
                username: '',
                lastLoginTime: '',
                lastLoginIp: ''
            },
            todoList: [],
            data: []
        };
    },
    components: {
        Schart
    },
    computed: {
        role() {
            return this.name === 'admin' ? '超级管理员' : '普通用户';
        }
    },
    created() {
        this.getUserInfo();
        this.getTodoList();
    },
    //自定义指令
    directives: {
        focus: {
            inserted: function (el) {
                el.querySelector('input').focus();
            }
        }
    },

    methods: {
        changeDate() {
            const now = new Date().getTime();
            this.data.forEach((item, index) => {
                const date = new Date(now - (6 - index) * 86400000);
                item.name = `${date.getFullYear()}/${date.getMonth() + 1}/${date.getDate()}`;
            });
        },
        getUserInfo() {
            request
                .get('/admin/get_user_info')
                .then((res) => {
                    if (res.code == 200) {
                        this.userInfo.username = res.data.username;
                        this.userInfo.lastLoginTime = res.data.last_login_time;
                        this.userInfo.lastLoginIp = res.data.last_login_ip;
                    } else {
                        this.$message.error(res.msg);
                        return false;
                    }
                })
                .catch((res) => {
                    console.log(res);
                });
        },
        getTodoList() {
            request
                .get('/admin/todolist')
                .then((res) => {
                    if (res.code == 200) {
                        this.todoList = [];
                        res.data.list.forEach((item) => {
                            item.showInput = false;
                            this.todoList.push(item);
                        });
                    } else {
                        this.$message.error(res.msg);
                        return false;
                    }
                })
                .catch((res) => {
                    console.log(res);
                });
        },
        handleSaveTodo(row) {
            row.showInput = !row.showInput;
            let url = row.id ? '/admin/update_todolist' : '/admin/add_todolist';
            if (row.todo != row.oldTodo) {
                request
                    .post(url, row)
                    .then((res) => {
                        if (res.code == 200) {
                            this.$message.success(res.msg);
                            this.getTodoList();
                        } else {
                            this.$message.error(res.msg);
                        }
                    })
                    .catch((res) => {
                        console.log(res);
                    });
            }
        },
        //点击输入框聚焦
        focusEvent(row) {
            row.oldTodo = row.todo;
        },

        addTodo() {
            if (this.todoList.length == 0 || this.todoList[this.todoList.length - 1]['id'] != 0) {
                this.todoList.push({
                    id: 0,
                    todo: '',
                    oldTodo: '',
                    finished_status: false,
                    showInput: true
                });
            }
        },
        handleEditTodo(row) {
            row.showInput = true;
        },
        handleDeleteTodo(row) {
            // 二次确认删除
            this.$confirm('确定要删除吗？', '提示', {
                type: 'warning'
            })
                .then(() => {
                    request
                        .delete('/admin/delete_todolist?id=' + row.id)
                        .then((res) => {
                            if (res.code == 200) {
                                this.$message.success('删除成功');
                                this.getTodoList();
                            } else {
                                this.$message.error(res.msg);
                                return false;
                            }
                        })
                        .catch((res) => {
                            console.log(res);
                        });
                })
                .catch(() => {});
        }
    }
};
</script>


<style scoped>
.el-row {
    margin-bottom: 20px;
}

.grid-content {
    display: flex;
    align-items: center;
    height: 100px;
}

.grid-cont-right {
    flex: 1;
    text-align: center;
    font-size: 14px;
    color: #999;
}

.grid-num {
    font-size: 30px;
    font-weight: bold;
}

.grid-con-icon {
    font-size: 50px;
    width: 100px;
    height: 100px;
    text-align: center;
    line-height: 100px;
    color: #fff;
}

.grid-con-1 .grid-con-icon {
    background: rgb(45, 140, 240);
}

.grid-con-1 .grid-num {
    color: rgb(45, 140, 240);
}

.grid-con-2 .grid-con-icon {
    background: rgb(100, 213, 114);
}

.grid-con-2 .grid-num {
    color: rgb(45, 140, 240);
}

.grid-con-3 .grid-con-icon {
    background: rgb(242, 94, 67);
}

.grid-con-3 .grid-num {
    color: rgb(242, 94, 67);
}

.user-info {
    display: flex;
    align-items: center;
    padding-bottom: 20px;
    border-bottom: 2px solid #ccc;
    margin-bottom: 20px;
}

.user-avator {
    width: 120px;
    height: 120px;
    border-radius: 50%;
}

.user-info-cont {
    padding-left: 50px;
    flex: 1;
    font-size: 14px;
    color: #999;
}

.user-info-cont div:first-child {
    font-size: 30px;
    color: #222;
}

.user-info-list {
    font-size: 14px;
    color: #999;
    line-height: 25px;
}

.user-info-list span {
    margin-left: 70px;
}

.mgb20 {
    margin-bottom: 20px;
}

.todo-item {
    font-size: 14px;
}

.todo-item-del {
    text-decoration: line-through;
    color: #999;
}

.schart {
    width: 100%;
    height: 300px;
}
</style>
