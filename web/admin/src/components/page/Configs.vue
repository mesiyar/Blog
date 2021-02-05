<template>
    <div>
        <div class="crumbs">
            <el-breadcrumb separator="/">
                <el-breadcrumb-item> <i class="el-icon-lx-cascades"></i> 配置管理 </el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <div class="container">
            <div class="handle-box">
                <!-- <el-button
                    type="primary"
                    icon="el-icon-delete"
                    class="handle-del mr10"
                    @click="delAllSelection"
                >批量删除</el-button> -->
                <el-input v-model="query.name" placeholder="关键字搜索" class="handle-input mr10"></el-input>
                <el-button type="primary" icon="el-icon-search" @click="handleSearch">搜索</el-button>
                <el-button type="success"  @click="handleCreate()">创建</el-button>
            </div>
            <el-table
                :data="tableData"
                border
                class="table"
                ref="multipleTable"
                header-cell-class-name="table-header"
                @selection-change="handleSelectionChange"
            >
                <el-table-column type="selection" width="55" align="center"></el-table-column>
                <el-table-column prop="id" label="ID" width="55" align="center"></el-table-column>
                <el-table-column prop="config_name" label="配置名"></el-table-column>
                <el-table-column prop="config_value" label="值"></el-table-column>
                <el-table-column label="操作" width="180" align="center">
                    <template slot-scope="scope">
                        <el-button type="text" icon="el-icon-edit" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
                        <el-button type="text" icon="el-icon-delete" class="red" @click="handleDelete(scope.$index, scope.row)"
                            >删除</el-button
                        >
                    </template>
                </el-table-column>
            </el-table>
            <div class="pagination">
                <el-pagination
                    background
                    layout="total, prev, pager, next"
                    :current-page="query.pageIndex"
                    :page-size="query.pageSize"
                    :total="pageTotal"
                    @current-change="handlePageChange"
                ></el-pagination>
            </div>
        </div>

        <!-- 编辑弹出框 -->
        <el-dialog title="编辑" :visible.sync="editVisible" width="80%">
            <el-form ref="form" :model="form" label-width="70px">
                <el-form-item label="配置名">
                    <el-input v-model="form.config_name"></el-input>
                </el-form-item>
                <el-form-item label="配置值">
                    <el-input v-model="form.config_value"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="editVisible = false">取 消</el-button>
                <el-button type="primary" @click="saveEdit">确 定</el-button>
            </span>
        </el-dialog>
        <!-- 创建弹出框 -->
        <el-dialog title="新建" :visible.sync="createVisible" width="80%" height="600px">
            <el-form ref="form" :model="createForm" label-width="70px">
                <el-form-item label="配置名">
                    <el-input v-model="createForm.config_name"></el-input>
                </el-form-item>
                <el-form-item label="配置值">
                    <el-input v-model="createForm.config_value"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="createVisible = false">取 消</el-button>
                <el-button type="primary" @click="saveCreate">确 定</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
import { fetchConfigs } from '../../api/index';
import request from '../../utils/request';
export default {
    name: 'basetable',
    data() {
        return {
            query: {
                address: '',
                name: '',
                pageIndex: 1,
                pageSize: 10
            },
            tableData: [],
            multipleSelection: [],
            delList: [],
            createVisible: false,
            editVisible: false,
            pageTotal: 0,
            form: {},
            createForm: {},
            idx: -1,
            id: -1,
            content: '',
            html: '',
            configs: {},
            tags: {}
        };
    },

    components: {
        
    },
    created() {
        this.getData();
    },
    methods: {
        
        getData() {
            fetchConfigs(this.query).then((res) => {
                console.log(res);
                this.tableData = res.data.lists;
                this.pageTotal = res.data.total;
            });
        },
        // 触发搜索按钮
        handleSearch() {
            this.$set(this.query, 'pageIndex', 1);
            this.getData();
        },
        // 删除操作
        handleDelete(index, row) {
            // 二次确认删除
            this.$confirm('确定要删除吗？', '提示', {
                type: 'warning'
            })
                .then(() => {
                    request
                        .delete('/admin/delete_config?id=' + row.id)
                        .then((res) => {
                            if (res.code == 200) {
                                this.$message.success('删除成功');
                                this.getData();
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
        },
        // 多选操作
        handleSelectionChange(val) {
            this.multipleSelection = val;
        },
        delAllSelection() {
            const length = this.multipleSelection.length;
            let str = '';
            this.delList = this.delList.concat(this.multipleSelection);
            for (let i = 0; i < length; i++) {
                str += this.multipleSelection[i].name + ' ';
            }
            this.$message.error(`删除了${str}`);
            this.multipleSelection = [];
        },
        // 创建操作
        handleCreate() {
            this.createVisible = true;
        },
        // 编辑操作
        handleEdit(index, row) {
            this.idx = index;
            this.form = row;
            console.log(row);
            this.editVisible = true;
        },
        // 保存编辑
        saveEdit() {
            this.editVisible = false;
            let fd = new FormData();
            fd.append('id', this.form.id);
            fd.append('config_name', this.form.config_name);
            fd.append('config_value', this.form.config_value);
            request
                .post('/admin/update_config', fd)
                .then((res) => {
                    if (res.code == 200) {
                        this.$message.success('编辑成功');
                        this.createVisible = false;
                        this.getData();
                    } else {
                        this.$message.error(res.msg);
                        return false;
                    }
                })
                .catch((res) => {
                    console.log(res);
                });
        },
        // 保存编辑
        saveCreate() {
            this.createVisible = false;
            console.log(this.createForm);
            let fd = new FormData();
            fd.append('config_name', this.createForm.config_name);
            fd.append('config_value', this.createForm.config_value);
            request
                .post('/admin/add_config', fd)
                .then((res) => {
                    if (res.code == 200) {
                        this.$message.success('创建成功');
                        this.createVisible = false;
                        this.getData();
                    } else {
                        this.$message.error(res.msg);
                        return false;
                    }
                })
                .catch((res) => {
                    console.log(res);
                });
        },
        // 分页导航
        handlePageChange(val) {
            this.$set(this.query, 'pageIndex', val);
            this.getData();
        },
        
        formatDate(row, column) {
            let date = new Date(parseInt(row.created_on) * 1000);
            let Y = date.getFullYear() + '-';
            let M = date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) + '-' : date.getMonth() + 1 + '-';
            let D = date.getDate() < 10 ? '0' + date.getDate() + ' ' : date.getDate() + ' ';
            let h = date.getHours() < 10 ? '0' + date.getHours() + ':' : date.getHours() + ':';
            let m = date.getMinutes() < 10 ? '0' + date.getMinutes() + ':' : date.getMinutes() + ':';
            let s = date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds();
            return Y + M + D + h + m + s;
        }
    }
};
</script>

<style scoped>
.handle-box {
    margin-bottom: 20px;
}

.handle-select {
    width: 120px;
}

.handle-input {
    width: 300px;
    display: inline-block;
}
.table {
    width: 100%;
    font-size: 14px;
}
.red {
    color: #ff0000;
}
.mr10 {
    margin-right: 10px;
}
.table-td-thumb {
    display: block;
    margin: auto;
    width: 40px;
    height: 40px;
}
.editor-btn {
    margin-top: 20px;
}
</style>
