import { defineConfig, presetAttributify, presetWind3 } from "unocss";

export default defineConfig({
  presets: [presetAttributify({}), presetWind3()],
  shortcuts: [
    // 常用布局样式
    ["flex-center", "flex items-center justify-center"],
    ["flex-between", "flex items-center justify-between"],
    [
      "absolute-center",
      "absolute top-0 left-0 right-0 bottom-0 flex items-center justify-center",
    ],

    // 常用间距
    [/^p-(\d+)$/, ([_, num]) => `p-${parseInt(num) * 4}`],
    [/^m-(\d+)$/, ([_, num]) => `m-${parseInt(num) * 4}`],

    // 常用颜色
    ["text-primary", "text-blue-500"],
    ["bg-primary", "bg-blue-500"],
    ["border-primary", "border-blue-500"],
  ],
  rules: [
    // 可以添加自定义规则
  ],
});
