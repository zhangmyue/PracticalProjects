package service

import (
	"testing"
)

// =========================================================================
// ================================ 单元测试 =================================
// =========================================================================

// 测试工具：断言函数
func assertEqual(t *testing.T, name string, expected, actual interface{}) {
	if expected != actual {
		t.Fatalf("%s 失败: 预期=%v, 实际=%v", name, expected, actual)
	}
}

func assertNotNil(t *testing.T, name string, err error) {
	if err == nil {
		t.Fatalf("%s 失败: 预期有错误，实际无错误", name)
	}
}

// 测试添加学生
func TestAddStudent(t *testing.T) {
	svc := NewStudentService()

	// 正常添加
	stu, err := svc.AddStudent("张三", 18)
	assertEqual(t, "AddStudent-err", nil, err)
	assertEqual(t, "AddStudent-ID", 1, stu.ID)
	assertEqual(t, "AddStudent-Name", "张三", stu.Name)

	// 名字为空
	_, err = svc.AddStudent("", 18)
	assertNotNil(t, "AddStudent-empty-name", err)

	// 年龄非法
	_, err = svc.AddStudent("李四", -5)
	assertNotNil(t, "AddStudent-invalid-age", err)
}

// 测试删除学生
func TestDeleteStudentByID(t *testing.T) {
	svc := NewStudentService()
	svc.AddStudent("张三", 18)

	// 删除存在的学生
	ok, err := svc.DeleteStudentByID(1)
	assertEqual(t, "Delete-ok", true, ok)
	assertEqual(t, "Delete-err", nil, err)

	// 删除不存在的学生
	ok, err = svc.DeleteStudentByID(999)
	assertEqual(t, "Delete-not-found-ok", false, ok)
	assertNotNil(t, "Delete-not-found-err", err)
}

// 测试更新学生
func TestUpdateStudentByID(t *testing.T) {
	svc := NewStudentService()
	svc.AddStudent("张三", 18)

	// 正常更新
	ok, err := svc.UpdateStudentByID(1, "张三三", 20)
	assertEqual(t, "Update-ok", true, ok)
	assertEqual(t, "Update-err", nil, err)

	stu, _ := svc.GetStudentByID(1)
	assertEqual(t, "Update-Name", "张三三", stu.Name)
	assertEqual(t, "Update-Age", 20, stu.Age)

	// 更新不存在的ID
	ok, err = svc.UpdateStudentByID(999, "test", 20)
	assertEqual(t, "Update-not-found", false, ok)
	assertNotNil(t, "Update-not-found-err", err)
}

// 测试查询单个学生
func TestGetStudentByID(t *testing.T) {
	svc := NewStudentService()
	svc.AddStudent("张三", 18)

	// 能查到
	stu, err := svc.GetStudentByID(1)
	assertEqual(t, "Get-err", nil, err)
	assertEqual(t, "Get-Name", "张三", stu.Name)

	// 查不到
	_, err = svc.GetStudentByID(999)
	assertNotNil(t, "Get-not-found", err)
}

// 测试列出所有学生
func TestListStudents(t *testing.T) {
	svc := NewStudentService()
	svc.AddStudent("张三", 18)
	svc.AddStudent("李四", 19)

	list, err := svc.ListStudents()
	assertEqual(t, "List-err", nil, err)
	assertEqual(t, "List-len", 2, len(list))
}

// =========================================================================
// ============================ Benchmark 性能测试 ===========================
// =========================================================================

// 性能测试：AddStudent
func BenchmarkAddStudent(b *testing.B) {
	svc := NewStudentService()
	// 重置计时器，避免初始化耗时影响结果
	b.ResetTimer()

	// b.N 是Go自动调整的循环次数
	for i := 0; i < b.N; i++ {
		_, _ = svc.AddStudent("Benchmark", 20)
	}
}

// 性能测试：GetStudentByID
func BenchmarkGetStudentByID(b *testing.B) {
	svc := NewStudentService()
	// 先插入100个学生，模拟真实数据
	for i := 0; i < 100; i++ {
		_, _ = svc.AddStudent("Student", 20)
	}

	b.ResetTimer()

	// 循环查询ID=50
	for i := 0; i < b.N; i++ {
		_, _ = svc.GetStudentByID(50)
	}
}

// 性能测试：ListStudents
func BenchmarkListStudents(b *testing.B) {
	svc := NewStudentService()
	// 先插入100个学生
	for i := 0; i < 100; i++ {
		_, _ = svc.AddStudent("Student", 20)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = svc.ListStudents()
	}
}
