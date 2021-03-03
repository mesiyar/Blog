<template>
    <div>
        <div class="crumbs">
            <el-breadcrumb separator="/">
                <el-breadcrumb-item><i class="el-icon-lx-cascades"></i> 文章管理</el-breadcrumb-item>
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
                <el-button type="success" @click="handleCreate()">创建</el-button>
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
                <el-table-column prop="title" label="标题"></el-table-column>
                <el-table-column prop="desc" label="描述"></el-table-column>
                <el-table-column prop="tag.name" label="标签"></el-table-column>
                <el-table-column label="状态" align="center">
                    <template slot-scope="scope">
                        <el-tag :type="scope.row.state == 1 ? 'success' : 'error'">{{ scope.row.state == 1 ? '正常' : '禁用'
                            }}
                        </el-tag>
                    </template>
                </el-table-column>
                <el-table-column label="发布状态" align="center">
                    <template slot-scope="scope">
                        <el-tag :type="scope.row.is_pub == 1 ? 'success' : 'error'">{{ scope.row.is_pub == 1 ? '发布' :
                            '草稿' }}
                        </el-tag>
                    </template>
                </el-table-column>
                <el-table-column label="置顶" align="center">
                    <template slot-scope="scope">
                        <el-tag :type="scope.row.is_top == 1 ? 'success' : 'error'">{{ scope.row.is_top == 1 ? '置顶' :
                            '无' }}
                        </el-tag>
                    </template>
                </el-table-column>

                <el-table-column prop="created_on" label="发表时间" :formatter="formatDate"></el-table-column>
                <el-table-column prop="read_nums" label="阅读量"></el-table-column>
                <el-table-column prop="comments" label="评论数"></el-table-column>
                <el-table-column label="操作" width="180" align="center">
                    <template slot-scope="scope">
                        <el-button type="text" icon="el-icon-edit" @click="handleEdit(scope.$index, scope.row)">编辑
                        </el-button>
                        <el-button type="text" icon="el-icon-delete" class="red"
                                   @click="handleDelete(scope.$index, scope.row)"
                        >删除
                        </el-button
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
                <el-form-item label="标题">
                    <el-input v-model="form.title"></el-input>
                </el-form-item>
                <el-form-item label="标签">
                    <el-select
                            v-model="form.tag_id"
                            placeholder="请选择标签"
                            clearable
                            :style="{ width: '200px' }"
                    >
                        <el-option
                                v-for="dict in tags"
                                :key="dict.id"
                                :label="dict.name"
                                :value="dict.id"
                        />
                    </el-select>
                </el-form-item>

                <el-form-item label="置顶">
                    <el-switch
                            v-model="form.is_top"
                            active-text="是"
                            active-value="1"
                            inactive-value="0"
                            inactive-text="否">
                    </el-switch>
                </el-form-item>
                <el-form-item label="草稿">
                    <el-switch
                            v-model="form.is_pub"
                            active-text="是"
                            active-value="1"
                            inactive-value="0"
                            inactive-text="否">
                    </el-switch>
                </el-form-item>

                <el-form-item label="描述">
                    <el-input v-model="form.desc"></el-input>
                </el-form-item>
                <el-form-item label="内容">
                    <mavon-editor v-model="form.content" ref="md" @imgAdd="$imgAdd" @change="change"
                                  style="min-height: 600px"/>
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
                <el-form-item label="标题">
                    <el-input v-model="createForm.title"></el-input>
                </el-form-item>
                <el-form-item label="标签">
                    <el-select
                            v-model="createForm.tag_id"
                            placeholder="请选择标签"
                            clearable
                            :style="{ width: '200px' }"
                    >
                        <el-option
                                v-for="dict in tags"
                                :key="dict.id"
                                :label="dict.name"
                                :value="dict.id"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="描述">
                    <el-input v-model="createForm.desc"></el-input>
                </el-form-item>

                <el-form-item label="置顶">
                    <el-switch
                            v-model="createForm.is_top"
                            active-text="是"
                            active-value="1"
                            inactive-value="0"
                            inactive-text="否">
                    </el-switch>
                </el-form-item>
                <el-form-item label="草稿">
                    <el-switch
                            v-model="createForm.is_pub"
                            active-text="是"
                            active-value="1"
                            inactive-value="0"
                            inactive-text="否">
                    </el-switch>
                </el-form-item>
                <el-form-item label="内容">
                    <mavon-editor v-model="createForm.content" ref="md" @imgAdd="$imgAdd" @change="change"
                                  style="min-height: 600px"/>
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
    import {fetchArticle} from '../../api/index';
    import {mavonEditor} from 'mavon-editor';
    import 'mavon-editor/dist/css/index.css';
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
            mavonEditor
        },
        created() {
            this.getData();
            this.getTags();
        },
        methods: {
            getTags() {
                request
                    .get('/admin/all_tags')
                    .then((res) => {
                        if (res.code == 200) {
                            this.tags = res.data.lists;
                        } else {
                            this.$message.error(res.msg);
                            return false;
                        }
                    })
                    .catch((res) => {
                        console.log(res);
                    });
            },

            getData() {
                fetchArticle(this.query).then((res) => {
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
                            .delete('/admin/delete_article?id=' + row.id)
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
                    .catch(() => {
                    });
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
                fd.append('title', this.form.title);
                fd.append('desc', this.form.desc);
                fd.append('content', this.form.content);
                fd.append('tag_id', this.form.tag_id);
                fd.append('is_top', this.form.is_top);
                fd.append('is_pub', this.form.is_pub);
                console.log(this.form);
                request
                    .post('/admin/update_article', fd)
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
                fd.append('title', this.createForm.title);
                fd.append('desc', this.createForm.desc);
                fd.append('content', this.createForm.content);
                fd.append('tag_id', this.createForm.tag_id);
                fd.append('is_top', this.createForm.is_top);
                fd.append('is_pub', this.createForm.is_pub);
                request
                    .post('/admin/add_article', fd)
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
            // 将图片上传到服务器，返回地址替换到md中
            $imgAdd(pos, $file) {
                var formdata = new FormData();
                formdata.append('file', $file);
                // 这里没有服务器供大家尝试，可将下面上传接口替换为你自己的服务器接口
                this.$axios({
                    url: '/common/upload',
                    method: 'post',
                    data: formdata,
                    headers: {'Content-Type': 'multipart/form-data'}
                }).then((url) => {
                    this.$refs.md.$img2Url(pos, url);
                });
            },
            change(value, render) {
                // render 为 markdown 解析后的结果
                this.html = render;
            },
            submit() {
                console.log(this.content);
                console.log(this.html);
                this.$message.success('提交成功！');
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
