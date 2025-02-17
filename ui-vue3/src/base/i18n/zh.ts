/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import type { I18nType } from './type.ts'

const words: I18nType = {
  destinationRuleDomain: {
    YAMLView: 'YAML视图',
    formView: '表单视图'
  },
  virtualServiceDomain: {
    YAMLView: 'YAML视图',
    formView: '表单视图'
  },
  dynamicConfigDomain: {
    YAMLView: 'YAML视图',
    formView: '表单视图',
    event: '事件'
  },
  routingRuleDomain: {
    YAMLView: 'YAML视图',
    formView: '表单视图'
  },
  addRoutingRuleDomain: {
    YAMLView: 'YAML视图',
    formView: '表单视图'
  },
  tagRuleDomain: {
    YAMLView: 'YAML视图',
    formView: '表单视图'
  },
  flowControlDomain: {
    notSet: '未设置',
    versionRecords: '版本记录',
    YAMLView: 'YAML视图',
    addConfiguration: '增加配置',
    addConfigurationItem: '增加配置项',
    addFilter: '增加筛选',
    configurationItem: '配置项',
    actuatingRange: '作用范围',
    scopeScreening: '作用范围筛选',
    endOfAction: '作用端',
    actions: '操作',
    filterType: '筛选类型',
    labelName: '标签名',
    formView: '表单视图',
    addMatch: '增加匹配',
    addRouter: '增加路由',
    addLabel: '增加标签',
    addressSubsetMatching: '地址子集匹配',
    value: '值',
    relation: '关系',
    requestParameterMatching: '请求参数匹配',
    matchingDimension: '匹配维度',
    parameter: '参数',
    ruleName: '规则名',
    actionObject: '作用对象',
    faultTolerantProtection: '容错保护',
    runTimeEffective: '运行时生效',
    ruleGranularity: '规则粒度',
    effectTime: '生效时间',
    enabledState: '启用状态',
    priority: '优先级',
    off: '关',
    on: '开',
    opened: '开启',
    closed: '关闭',
    enabled: '启用',
    disabled: '禁用'
  },
  instanceDomain: {
    flowDisabled: '流量禁用',
    operatorLog: '执行日志',
    CPU: 'CPU',
    enableAppInstanceLogs: '开启该应用所有实例的访问日志',
    appServiceLoadBalance: '调整应用提供服务的负载均衡策略',
    appServiceNegativeClusteringMethod: '调整应用提供服务的的集群方式',
    appServiceRetries: '调整该应用提供服务的重试次数',
    appServiceTimeout: '调整应用提供服务的超时时间',
    close: '关闭',
    enable: '开启',
    executionLog: '执行日志',
    retryCount: '重试次数',
    clusterApproach: '集群方式',
    timeout_ms: '超时时间(ms)',
    details: '详情',
    loadBalance: '负载均衡',
    monitor: '监控',
    linkTracking: '链路追踪',
    configuration: '场景配置',
    event: '事件',
    healthExamination_k8s: '健康检查(k8s)',
    instanceLabel: '实例标签',
    instanceImage_k8s: '镜像(k8s)',
    owningWorkload_k8s: '所属工作负载(k8s)',
    node: '节点',
    whichApplication: '所属应用',
    registerCluster: '注册集群',
    dubboPort: 'Dubbo端口',
    instanceIP: '实例IP',
    ip: 'IP',
    name: '实例名称',
    deployState: '部署状态',
    deployCluster: '部署集群',
    deployClusters: '部署集群',
    registerState: '注册状态',
    registerClusters: '注册集群',
    registryClusters: '注册集群',
    cpu: 'CPU',
    memory: '内存',
    startTime: '启动时间',
    registerTime: '注册时间',
    labels: '标签',
    instanceCount: '实例数量',
    instanceName: '实例名称',
    creationTime_k8s: '创建时间(k8s)',
    startTime_k8s: '启动时间(k8s)'
  },
  serviceDomain: {
    name: '服务名',
    notSpecified: '不指定'
  },
  service: '服务',
  versionGroup: '版本&分组',
  avgQPS: 'QPS',
  avgRT: 'RT',
  requestTotal: '近1min请求总量',
  serviceSearch: '服务查询',
  serviceGovernance: '路由规则',
  trafficManagement: '流量管控',
  serviceMetrics: '服务统计',
  serviceRelation: '服务关系',
  routingRule: '条件路由',
  tagRule: '标签路由',
  meshRule: 'Mesh路由',
  dynamicConfig: '动态配置',
  accessControl: '黑白名单',
  weightAdjust: '权重调整',
  loadBalance: '负载均衡',
  serviceTest: '服务测试',
  serviceMock: '服务Mock',
  services: '服务',
  providers: '提供者',
  consumers: '消费者',
  application: '应用',
  instance: '实例',
  all: '全部',
  common: '通用',

  metrics: '可观测',
  relation: '关系',
  group: '组',
  version: '版本',
  app: '应用',
  ip: 'IP地址',
  qps: 'qps',
  rt: 'rt',
  successRate: '成功率',
  serviceInfo: '服务信息',
  port: '端口',
  timeout: '超时(毫秒)',
  serialization: '序列化',
  appName: '应用名',
  instanceNum: '实例数量',
  deployCluster: '部署集群',
  registerClusters: '注册集群列表',
  serviceName: '服务名',
  registrySource: '注册来源',
  instanceRegistry: '应用级',
  interfaceRegistry: '接口级',
  allRegistry: '应用级/接口级',
  operation: '操作',
  searchResult: '查询结果',
  search: '搜索',
  methodName: '方法名',
  enabled: '开启',
  disabled: '禁用',
  method: '方法',
  weight: '权重',
  create: '创建',
  save: '保存',
  cancel: '取消',
  close: '关闭',
  confirm: '确认',
  ruleContent: '规则内容',
  createNewRoutingRule: '创建新路由规则',
  createNewTagRule: '创建新标签规则',
  createMeshTagRule: '创建新mesh规则',
  createNewDynamicConfigRule: '创建新动态配置规则',
  createNewWeightRule: '新建权重规则',
  createNewLoadBalanceRule: '新建负载均衡规则',
  createTimeoutRule: '创建超时时间规则',
  createRetryRule: '创建重试规则',
  createRegionRule: '创建同区域优先规则',
  createArgumentRule: '创建参数路由规则',
  createMockCircuitRule: '创建调用降级规则',
  createAccesslogRule: '创建访问日志规则',
  createGrayRule: '创建灰度隔离规则',
  createWeightRule: '创建权重比例规则',
  serviceIdHint: '服务名',
  view: '查看',
  edit: '编辑',
  delete: '删除',
  searchRoutingRule: '搜索路由规则',
  searchAccessRule: '搜索黑白名单',
  searchWeightRule: '搜索权重调整规则',
  dataIdClassHint: '服务接口的类完整包路径',
  dataIdVersionHint: '服务接口的Version,根据接口实际情况选填',
  dataIdGroupHint: '服务接口的Group,根据接口实际情况选填',
  agree: '同意',
  disagree: '不同意',
  searchDynamicConfig: '搜索动态配置',
  appNameHint: '服务所属的应用名称',
  basicInfo: '基础信息',
  metaData: '元数据',
  methodMetrics: '服务方法统计',
  searchDubboService: '搜索Dubbo服务或应用',
  serviceSearchHint: '服务ID, org.apache.dubbo.demo.api.DemoService, * 代表所有服务',
  ipSearchHint: '在指定的IP地址上查找目标服务器提供的所有服务',
  appSearchHint: '输入应用名称以查找由一个特定应用提供的所有服务, * 代表所有',
  searchTagRule: '根据应用名搜索标签规则',
  searchMeshRule: '根据应用名搜索mesh规则',
  searchSingleMetrics: '输入IP搜索Metrics信息',
  searchBalanceRule: '搜索负载均衡规则',
  parameterList: '参数列表',
  returnType: '返回值',
  noMetadataHint:
    '无元数据信息，请升级至Dubbo2.7及以上版本，或者查看application.properties中关于config center的配置，详见',
  here: '这里',
  configAddress:
    'https://github.com/apache/incubator-dubbo-admin/wiki/Dubbo-Admin%E9%85%8D%E7%BD%AE%E8%AF%B4%E6%98%8E',
  whiteList: '白名单',
  whiteListHint: '白名单IP列表, 多个地址用逗号分隔: 1.1.1.1,2.2.2.2',
  blackList: '黑名单',
  blackListHint: '黑名单IP列表, 多个地址用逗号分隔: 3.3.3.3,4.4.4.4',
  address: '地址列表',
  weightAddressHint: '此权重设置的IP地址,用逗号分隔: 1.1.1.1,2.2.2.2',
  weightHint: '权重值，默认100',
  methodHint: '负载均衡生效的方法，*代表所有方法',
  strategy: '策略',
  balanceStrategyHint: '负载均衡策略',
  goIndex: '返回首页',
  releaseLater: '在后续版本中发布，敬请期待',
  later: {
    metrics: 'Metrics会在后续版本中发布，敬请期待',
    serviceTest: '服务测试会在后续版本中发布，敬请期待',
    serviceMock: '服务Mock会在后续版本中发布，敬请期待'
  },
  by: '按',
  $vuetify: {
    dataIterator: {
      rowsPerPageText: '每页记录数：',
      rowsPerPageAll: '全部',
      pageText: '{0}-{1} 共 {2} 条',
      noResultsText: '没有找到匹配记录',
      nextPage: '下一页',
      prevPage: '上一页'
    },
    dataTable: {
      rowsPerPageText: '每页行数：'
    },
    noDataText: '无可用数据'
  },
  configManage: '配置管理',
  configCenterAddress: '配置中心地址',
  searchDubboConfig: '搜索Dubbo配置',
  createNewDubboConfig: '新建Dubbo配置',
  scope: '范围',
  name: '名称',
  warnDeleteConfig: ' 是否要删除Dubbo配置: ',
  warnDeleteRouteRule: '是否要删除路由规则',
  warnDeleteDynamicConfig: '是否要删除动态配置',
  warnDeleteBalancing: '是否要删除负载均衡规则',
  warnDeleteAccessControl: '是否要删除黑白名单',
  warnDeleteTagRule: '是否要删除标签路由',
  warnDeleteMeshRule: '是否要删除mesh路由',
  warnDeleteWeightAdjust: '是否要删除权重规则',
  configNameHint: '配置所属的应用名, global 表示全局配置',
  configContent: '配置内容',
  testMethod: '测试方法',
  execute: '执行',
  result: '结果: ',
  success: ' 成功',
  fail: '失败',
  detail: '详情',
  more: '更多',
  copyUrl: '复制 URL',
  copy: '复制',
  url: 'URL',
  copySuccessfully: '已复制',
  test: '测试',
  placeholders: {
    searchService: '通过服务名搜索服务'
  },
  methods: '方法列表',
  testModule: {
    searchServiceHint: '完整服务ID, org.apache.dubbo.demo.api.DemoService, 按回车键查询'
  },
  userName: '用户名',
  password: '密码',
  login: '登录',
  apiDocs: '接口文档',
  apiDocsRes: {
    dubboProviderIP: 'Dubbo 提供者Ip',
    dubboProviderPort: 'Dubbo 提供者端口',
    loadApiList: '加载接口列表',
    apiListText: '接口列表',
    apiForm: {
      missingInterfaceInfo: '缺少接口信息',
      getApiInfoErr: '获取接口信息异常',
      api404Err: '接口名称不正确,没有查找到接口参数和响应信息',
      apiRespDecShowLabel: '响应说明',
      apiNameShowLabel: '接口名称',
      apiPathShowLabel: '接口位置',
      apiMethodParamInfoLabel: '接口参数',
      apiVersionShowLabel: '接口版本',
      apiGroupShowLabel: '接口分组',
      apiDescriptionShowLabel: '接口说明',
      isAsyncFormLabel: '是否异步调用(此参数不可修改,根据接口定义的是否异步显示)',
      apiModuleFormLabel: '接口模块(此参数不可修改)',
      apiFunctionNameFormLabel: '接口方法名(此参数不可修改)',
      registryCenterUrlFormLabel: '注册中心地址, 如果为空将使用Dubbo 提供者Ip和端口进行直连',
      paramNameLabel: '参数名',
      paramPathLabel: '参数位置',
      paramDescriptionLabel: '说明',
      paramRequiredLabel: '该参数为必填',
      doTestBtn: '测试',
      responseLabel: '响应',
      responseExampleLabel: '响应示例',
      apiResponseLabel: '接口响应',
      LoadingLabel: '加载中...',
      requireTip: '有未填写的必填项',
      requireItemTip: '该项为必填!',
      requestApiErrorTip: '请求接口发生异常,请检查提交的数据,特别是JSON类数据和其中的枚举部分',
      unsupportedHtmlTypeTip: '暂不支持的表单类型',
      none: '无'
    }
  },
  authFailed: '权限验证失败',

  ruleList: '规则列表',
  mockRule: '规则配置',
  mockData: '模拟数据',
  globalDisable: '全局禁用',
  globalEnable: '全局启用',
  saveRuleSuccess: '保存规则成功',
  deleteRuleSuccess: '删除成功',
  disableRuleSuccess: '禁用成功',
  enableRuleSuccess: '启用成功',
  methodNameHint: '服务方法名',
  createMockRule: '创建规则',
  editMockRule: '修改规则',
  deleteRuleTitle: '确定要删除此服务Mock规则吗？',

  ruleName: '规则名',
  ruleGranularity: '规则粒度',
  createTime: '创建时间',
  lastModifiedTime: '最后修改时间',
  enable: '是否启用',
  protection: '容错保护',
  trafficTimeout: '超时时间',
  trafficRetry: '调用重试',
  trafficRegion: '同区域优先',
  trafficIsolation: '环境隔离',
  trafficWeight: '权重比例',
  trafficArguments: '参数路由',
  trafficMock: '调用降级',
  trafficAccesslog: '访问日志',
  trafficHost: '固定机器导流',
  trafficGray: '流量灰度',

  homePage: '集群概览',
  serviceManagement: '开发测试',

  groupInputPrompt: '请输入服务group(可选)',
  versionInputPrompt: '请输入服务version(可选)',
  resources: '资源详情',
  applications: '应用',
  instances: '实例',
  applicationDomain: {
    instanceCount: '实例数量',
    deployClusters: '部署集群',
    registryClusters: '注册集群',
    registerClusters: '注册集群',
    registerModes: '注册模式',
    operatorLog: '执行日志',
    flowWeight: '流量权重',
    gray: '灰度隔离',
    name: '应用名',
    detail: '详情',
    instance: '实例',
    service: '服务',
    monitor: '监控',
    tracing: '链路追踪',
    config: '配置',
    event: '事件',
    appName: '应用名',
    rpcProtocols: 'RPC 协议',
    dubboVersions: 'Dubbo 版本',
    dubboPorts: 'Dubbo 端口',
    serialProtocols: '序列化协议',
    appTypes: '应用类型',
    images: '应用镜像',
    workloads: '工作负载',
    deployCluster: '部署集群',
    registerCluster: '注册集群',
    registerMode: '注册模式'
  },

  searchDomain: {
    total: '共计',
    unit: '条'
  },
  messageDomain: {
    success: {
      copy: '您已经成功复制一条信息'
    }
  },
  backHome: '回到首页',
  noPageTip: '抱歉，你访问的页面不存在',
  globalSearchTip: '搜索ip，应用，实例，服务',

  placeholder: {
    typeAppName: '请输入应用名，支持前缀搜索',
    typeDefault: '请输入',
    typeRoutingRules: '搜索路由规则，支持前缀过滤'
  },
  none: '无',
  details: '详情',
  debug: '调试',
  distribution: '分布',
  monitor: '监控',
  tracing: '链路追踪',
  sceneConfig: '场景配置',
  event: '事件',

  provideService: '提供服务',
  dependentService: '依赖服务',
  idx: '序号',
  submit: '提交',
  reset: '重置',
  router: {
    resource: {
      app: {
        list: '列表'
      },
      ins: {
        list: '列表'
      },
      svc: {
        list: '列表'
      }
    }
  }
}

export default words
