# ifndef uint64_t
#  if  __WORDSIZE == 64
typedef unsigned long long uint64_t;
#  else
__extension__
typedef unsigned long long int uint64_t;
#  endif
# endif

typedef struct __qrng_device_t__ qrng_device_t; /* 打开设备后返回的设备句柄 */
#define QRNG_MAX_READ_SIZE (1000000000) /* 单次读取最大长度， 单位字节 */

/* the qrng max devs support*/
#define QRNG_MAX_QRNG_DEVS (256)

typedef struct __qrng_dev_info_t__ {
    char sn[64];
    unsigned int dev_version;
    unsigned int dev_status;
    float rand_rate;
    uint64_t total_rand;
} qrng_dev_info_t;

typedef struct __qrng_all_info_t__ {
    unsigned int lib_version;
    unsigned int driver_version;
    unsigned int total_devs;
    qrng_dev_info_t qrng_info[QRNG_MAX_QRNG_DEVS];
} qrng_all_info_t;

// 获取设备数量接口
unsigned int qrng_get_count(void);

// 根据设备编号获取设备随机数当前速率
int qrng_get_rand_rate(unsigned int device_num, float *rand_rate);

// 根据设备编号获取设备随机数生成总量
int qrng_get_total_rand(unsigned int device_num, uint64_t *total_num);

// 打开设备接口
int qrng_dev_open(unsigned int device_num, qrng_device_t **pthiz);

// 关闭设备
void qrng_dev_close(qrng_device_t *thiz);

// 根据设备句柄获取随机数
int qrng_read_random(qrng_device_t *thiz, void *buffer, unsigned int size);

//根据设备编号使能设备
int qrng_dev_enable(unsigned int device_num);

// 根据设备编号禁止设备
int qrng_dev_disable(unsigned int device_num);

// 获取 QRNG 设备版本号
int qrng_get_version(unsigned int device_num, unsigned int *lib_version,
                     unsigned int *driver_version, unsigned int *dev_version);

// 根据设备编号复位设备
int qrng_dev_reset(unsigned int device_num);

// 读取当前 QRNG 的设备状态
int qrng_get_status(unsigned int device_num, unsigned int *status);

//根据设备句柄获取熵源数据
int qrng_read_entropy(qrng_device_t *thiz, void *buffer, unsigned int size);

// 根据设备编号获取 QRNG 设备序列号
int qrng_get_sn(unsigned int device_num, char *sn);

// 获取当前建议调度的 QRNG 设备编号
int qrng_sched_number(void);

// 根据设备编号获取随机数
int qrng_read_random_number(unsigned int device_num, void *buffer,
                            unsigned int size);

// 根据当前建议的设备编号获取随机数
int qrng_read_random_sched(void *buffer, unsigned int size);

// 获取错误码详细说明
char *qrng_get_str_error(int errNumber);

// 获取所有 QRNG 设备信息
int qrng_get_all_info(qrng_all_info_t *all_info);
