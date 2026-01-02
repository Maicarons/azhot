# azhot

<p align="center">
  <img src="../banner.jpg" alt="Banner" style="max-width:100%;height:auto;" />
</p>

azhot 是一个聚合各大平台热搜数据的前端项目，提供统一的界面浏览各大平台的实时热搜榜单。

## 项目介绍

azhot 前端是一个基于 Vue 3 和 Element Plus 构建的热搜数据展示平台，旨在为用户提供各大互联网平台的实时热搜榜单浏览服务。

- 使用 Vue 3 的组合式 API 提升开发效率
- 基于 Element Plus 实现美观统一的 UI 风格
- 利用 Vite 实现快速热更新和构建
- 使用 UnoCSS 提供原子化CSS，提高样式开发效率

## 功能特性

- 实时查看超过30个主流平台的热搜数据
- 平台分类展示与详情浏览
- 历史查询记录功能
- 响应式布局适配移动端和桌面端
- 暗色模式支持

## 技术栈

- Vue 3.5.26
- Element Plus 2.13.0
- Vite 7.3.0
- TypeScript 5.0.0
- UnoCSS 0.61.0
- Axios 1.13.2
- vue-router 4.6.4

## 快速开始

1. 安装依赖：

```bash
pnpm install
```

2. 启动开发服务器：

```bash
pnpm dev
```

3. 构建生产版本：

```bash
pnpm build
```

4. 预览生产构建结果：

```bash
pnpm preview
```

## UnoCSS 配置

项目使用 UnoCSS 提供原子化 CSS 功能，配置文件位于 `uno.config.ts`，包含常用的布局和样式快捷方式。

## 项目结构

```
src/
├── components/
│   ├── Footer.vue
│   ├── Header.vue
│   ├── HistoryQuery.vue
│   ├── Home.vue
│   ├── PlatformDetail.vue
│   └── PlatformList.vue
├── types/
│   └── env.d.ts
├── App.vue
└── main.ts
```

## 贡献

欢迎提交 Issue 和 Pull Request 来帮助我们改进项目。

## 许可证

ISC