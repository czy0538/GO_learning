package mlib

import "errors"

type Music struct { //音乐类型定义
	Id     string //音乐ID
	Name   string //音乐名称
	Artist string //艺术家
	Source string //播放路径
	Type   string //音乐类型
}
type MusicManager struct { //音乐库类型
	musics []Music //成员为一个音乐类型的切片
}

//构造函数，初始化为0首音乐。构造函数的样式为NewXXX()
func NewMusicManager() *MusicManager {
	return &MusicManager{make([]Music, 0)} //初始音乐切片长度为0
}

//乐库类的Len方法，获取乐库的音乐数量
func (m *MusicManager) Len() int {
	return len(m.musics) //返回切片的长度即歌曲的数量
}

//乐库类的Get方法,获取索引值为index的音乐信息
func (m *MusicManager) Get(index int) (music *Music, err error) {
	if index < 0 || index >= len(m.musics) { //索引不在合理范围
		return nil, errors.New("Index out of range.") //返回nil和错误信息
	}
	return &m.musics[index], nil //返回曲库中索引为index的音乐信息和nil
}

//乐库类的按音乐名称查找方法
func (m *MusicManager) Find(name string) (*Music, int) {
	if len(m.musics) == 0 { //乐库为空
		return nil, -1 //返回nil,-1
	}
	for i, m := range m.musics { //乐库不为空，遍历音乐库
		if m.Name == name { //找到
			return &m, i
		}
	}
	return nil, -1 //乐库不为空，但没找到返回nil,-1
}

//乐库类的增加音乐方法
func (m *MusicManager) Add(music *Music) {
	m.musics = append(m.musics, *music) //追加
}

//乐库类的删除方法,
func (m *MusicManager) Remove(index int) *Music {
	if index < 0 || index >= len(m.musics) {
		return nil //删除的索引值范围不对，返回空
	}
	removeMusic := m.musics[index] //找到要删除的音乐信息赋给removeMusic
	//从切片中删除元素
	if index < len(m.musics)-1 { //中间元素
		m.musics = append(m.musics[:index], m.musics[index+1:]...)
	} else if index == 0 { //删除仅有的一个元素
		m.musics = make([]Music, 0)
	} else { //删除最后一个元素
		m.musics = m.musics[:index-1]
	}
	return &removeMusic //返回被删除的音乐信息
}
