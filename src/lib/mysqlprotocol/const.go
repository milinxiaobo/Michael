package mysqlprotocol

import (
	"errors"
)

var (
	errNotFount       = errors.New("not found")
	errPacketTooShort = errors.New("packet too short")
	errPacketNotParse = errors.New("packet not parse")
	errPacketNotError = errors.New("packet not error") // 当前包还有剩余数据
	errPacketParsed   = errors.New("packet parsed")    // 当前包没有剩余数据
)

var serverStatusFlag = map[uint16]string{
	1:       "SERVER_STATUS_IN_TRANS",     //A transaction is currently active
	2:       "SERVER_STATUS_AUTOCOMMIT",   // Autocommit mode is set
	8:       "SERVER_MORE_RESULTS_EXISTS", //more results exists (more packet follow)
	16:      "SERVER_QUERY_NO_GOOD_INDEX_USED",
	32:      "SERVER_QUERY_NO_INDEX_USED",
	64:      "SERVER_STATUS_CURSOR_EXISTS",        // when using COM_STMT_FETCH, indicate that current cursor still has result (deprecated)
	128:     "SERVER_STATUS_LAST_ROW_SENT",        //	when using COM_STMT_FETCH, indicate that current cursor has finished to send results (deprecated)
	1 << 8:  "SERVER_STATUS_DB_DROPPED",           //	database has been dropped
	1 << 9:  "SERVER_STATUS_NO_BACKSLASH_ESCAPES", //	current escape mode is "no backslash escape"
	1 << 10: "SERVER_STATUS_METADATA_CHANGED",     //	A DDL change did have an impact on an existing PREPARE (an automatic reprepare has been executed)
	1 << 11: "SERVER_QUERY_WAS_SLOW",              //
	1 << 12: "SERVER_PS_OUT_PARAMs",               //	this resultset contain stored procedure output parameter
	1 << 13: "SERVER_STATUS_IN_TRANS_READONLY",    //	current transaction is a read-only transaction
	1 << 14: "SERVER_SESSION_STATE_CHANGED",       //	session state change. see Session change type for more information
}

var cmdMap = map[uint16]string{
	0x00: "COM_SLEEP",               //	（内部线程状态）	（无）
	0x01: "COM_QUIT",                //	关闭连接	mysql_close
	0x02: "COM_INIT_DB",             //	切换数据库	mysql_select_db
	0x03: "COM_QUERY",               //	SQL查询请求	mysql_real_query
	0x04: "COM_FIELD_LIST",          //	获取数据表字段信息	mysql_list_fields
	0x05: "COM_CREATE_DB",           //	创建数据库	mysql_create_db
	0x06: "COM_DROP_DB",             //	删除数据库	mysql_drop_db
	0x07: "COM_REFRESH",             //	清除缓存	mysql_refresh
	0x08: "COM_SHUTDOWN",            //	停止服务器	mysql_shutdown
	0x09: "COM_STATISTICS",          //	获取服务器统计信息	mysql_stat
	0x0A: "COM_PROCESS_INFO",        //	获取当前连接的列表	mysql_list_processes
	0x0B: "COM_CONNECT",             //	（内部线程状态）	（无）
	0x0C: "COM_PROCESS_KILL",        //	中断某个连接	mysql_kill
	0x0D: "COM_DEBUG",               //	保存服务器调试信息	mysql_dump_debug_info
	0x0E: "COM_PING",                //	测试连通性	mysql_ping
	0x0F: "COM_TIME",                //	（内部线程状态）	（无）
	0x10: "COM_DELAYED_INSERT",      //	（内部线程状态）	（无）
	0x11: "COM_CHANGE_USER",         //	重新登陆（不断连接）	mysql_change_user
	0x12: "COM_BINLOG_DUMP",         //	获取二进制日志信息	（无）
	0x13: "COM_TABLE_DUMP",          //	获取数据表结构信息	（无）
	0x14: "COM_CONNECT_OUT",         //	（内部线程状态）	（无）
	0x15: "COM_REGISTER_SLAVE",      //	从服务器向主服务器进行注册	（无）
	0x16: "COM_STMT_PREPARE",        //	预处理SQL语句	mysql_stmt_prepare
	0x17: "COM_STMT_EXECUTE",        //	执行预处理语句	mysql_stmt_execute
	0x18: "COM_STMT_SEND_LONG_DATA", //	发送BLOB类型的数据	mysql_stmt_send_long_data
	0x19: "COM_STMT_CLOSE",          //	销毁预处理语句	mysql_stmt_close
	0x1A: "COM_STMT_RESET",          //	清除预处理语句参数缓存	mysql_stmt_reset
	0x1B: "COM_SET_OPTION",          //	设置语句选项	mysql_set_server_option
	0x1C: "COM_STMT_FETCH",          //	获取预处理语句的执行结果	mysql_stmt_fetch
}

var comStmtExecuteFlag = map[uint16]string{
	0: "no cursor",
	1: "read only",
	2: "cursor for update",
	4: "scrollable cursor",
}

var resultSetFieldTypes = map[uint16]string{
	0:   "MYSQL_TYPE_DECIMAL",
	1:   "MYSQL_TYPE_TINY",
	2:   "MYSQL_TYPE_SHORT",
	3:   "MYSQL_TYPE_LONG",
	4:   "MYSQL_TYPE_FLOAT",
	5:   "MYSQL_TYPE_DOUBLE",
	6:   "MYSQL_TYPE_NULL",
	7:   "MYSQL_TYPE_TIMESTAMP",
	8:   "MYSQL_TYPE_LONGLONG",
	9:   "MYSQL_TYPE_INT24",
	10:  "MYSQL_TYPE_DATE",
	11:  "MYSQL_TYPE_TIME",
	12:  "MYSQL_TYPE_DATETIME",
	13:  "MYSQL_TYPE_YEAR",
	14:  "MYSQL_TYPE_NEWDATE",
	15:  "MYSQL_TYPE_VARCHAR",
	16:  "MYSQL_TYPE_BIT",
	17:  "MYSQL_TYPE_TIMESTAMP2",
	18:  "MYSQL_TYPE_DATETIME2",
	19:  "MYSQL_TYPE_TIME2",
	246: "MYSQL_TYPE_NEWDECIMAL",
	247: "MYSQL_TYPE_ENUM",
	248: "MYSQL_TYPE_SET",
	249: "MYSQL_TYPE_TINY_BLOB",
	250: "MYSQL_TYPE_MEDIUM_BLOB",
	251: "MYSQL_TYPE_LONG_BLOB",
	252: "MYSQL_TYPE_BLOB",
	253: "MYSQL_TYPE_VAR_STRING",
	254: "MYSQL_TYPE_STRING",
	255: "MYSQL_TYPE_GEOMETRY",
}

var resultSetFieldDetailFlag = map[uint16]string{
	1:    "NOT_NULL",         //field cannot be null
	2:    "PRIMARY_KEY",      //field is a primary key
	4:    "UNIQUE_KEY",       //field is unique
	8:    "MULTIPLE_KEY",     //field is in a multiple key
	16:   "BLOB",             //is this field a Blob
	32:   "UNSIGNED",         //is this field unsigned
	64:   "DECIMAL",          //is this field a decimal
	128:  "BINARY_COLLATION", //whether this field has a binary collation
	256:  "ENUM",             //Field is an enumeration
	512:  "AUTO_INCREMENT",   //field auto-increment
	1024: "TIMESTAMP",        //field is a timestamp value
	2048: "SET",              //field is a SET
}
