 --  associated_subtask 关联子任务 --   
CREATE TABLE `associated_subtask`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `task_id` bigint(0) UNSIGNED NULL DEFAULT NULL,
  `task_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `task_name_eng` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `task_details` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `task_details_eng` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `task_status` bigint(0) NULL DEFAULT NULL,
  `reward` bigint(0) NULL DEFAULT NULL,
  `experience` bigint(0) NULL DEFAULT NULL,
  `number` bigint(0) NULL DEFAULT NULL,
  `article` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `link` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `label` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `treasure_id` bigint(0) UNSIGNED NOT NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_associated_subtask_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_associated_subtask_treasure_id`(`treasure_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

 --  chest_collections 宝箱领取度 --  
CREATE TABLE `chest_collections`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `user_id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT NULL,
  `chest_amount` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_chest_collections_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_chest_collections_user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;


 --  daily_task 日常任务完成表（一日一更新） --  
CREATE TABLE `daily_task`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `user_id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT NULL,
  `task_id` bigint(0) UNSIGNED NOT NULL DEFAULT NULL,
  `complete` bigint(0) NULL DEFAULT NULL,
  `experience` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_daily_task_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_daily_task_user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;


 --  daily_task_bak 日常任务完成表(冷表,存储30天) --  
CREATE TABLE `daily_task_bak`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `user_id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `task_id` bigint(0) UNSIGNED NULL DEFAULT NULL,
  `complete` bigint(0) NULL DEFAULT NULL,
  `experience` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_daily_task_bak_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;


 --  label 标签 --  
CREATE TABLE `label`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `creator` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `content` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_label_creator`(`creator`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;


 --  participant 策展任务参与者列表 --  
CREATE TABLE `participant`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `user_id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT NULL,
  `user_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `nick_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `avatar` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `award_amount` decimal(30, 2) NOT NULL,
  `task_id` bigint(0) UNSIGNED NOT NULL DEFAULT NULL,
  `status` bigint(0) NULL DEFAULT 1,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_participant_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_participant_user_id`(`user_id`) USING BTREE,
  INDEX `idx_participant_task_id`(`task_id`) USING BTREE,
  INDEX `idx_participant_status`(`status`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;


 --  publish_task 策展任务表 --  
CREATE TABLE `publish_task`  (
        `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
        `created_at` datetime(0) NULL DEFAULT NULL,
        `updated_at` datetime(0) NULL DEFAULT NULL,
        `deleted_at` datetime(0) NULL DEFAULT NULL,
        `creator` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
        `status` bigint(0) NOT NULL DEFAULT 1,
        `tweet_address` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
        `label` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
        `award_budget` decimal(30, 2) NOT NULL,
        `max_user` bigint(0) NOT NULL,
        `award_amount` decimal(30, 2) NOT NULL,
        `end_time` datetime(0) NOT NULL,
        `accomplish` bigint(0) NOT NULL DEFAULT 0,
        PRIMARY KEY (`id`) USING BTREE,
        INDEX `idx_publish_task_tweet_address`(`tweet_address`) USING BTREE,
        INDEX `idx_publish_task_accomplish`(`accomplish`) USING BTREE,
        INDEX `idx_publish_task_deleted_at`(`deleted_at`) USING BTREE,
        INDEX `idx_publish_task_creator`(`creator`) USING BTREE,
        INDEX `idx_publish_task_status`(`status`) USING BTREE
        ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

 --  subtask_style 子任务样式表 --  
CREATE TABLE `subtask_style`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `task_id` bigint(0) UNSIGNED NULL DEFAULT NULL DEFAULT NULL,
  `task_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `task_name_eng` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `task_details` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `task_details_eng` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `task_status` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_subtask_style_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;


 --  task_demand 任务要求 --  
CREATE TABLE `task_demand`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `task_id` bigint(0) UNSIGNED NULL DEFAULT NULL,
  `task_name` bigint(0) NULL DEFAULT NULL,
  `task_demand` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `article` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_task_demand_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_task_demand_task_id`(`task_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;


 --  task_template 任务要求模板 --  
CREATE TABLE `task_template` (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `task_id` bigint(0) NULL DEFAULT NULL,
  `task_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

 --  treasure_stages 参与者任务要求完成度 --  
CREATE TABLE `treasure_stages`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `user_id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT NULL,
  `task_id` bigint(0) UNSIGNED NOT NULL DEFAULT NULL,
  `task_name` bigint(0) NULL DEFAULT NULL,
  `task_status` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_treasure_stages_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_treasure_stages_user_id`(`user_id`) USING BTREE,
  INDEX `idx_treasure_stages_task_id`(`task_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;


 --  treasure_task 宝箱样式 --  
CREATE TABLE `treasure_task`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `demand_integral` bigint(0) NOT NULL,
  `task_reward` bigint(0) NOT NULL,
  `experience_reward` bigint(0) NOT NULL,
  `reward_quantity` bigint(0) NOT NULL,
  `employ` bigint(0) NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_treasure_task_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

 --  treasure_task_stage 宝箱阶段  --  
CREATE TABLE `treasure_task_stage` (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `treasure` bigint(0) UNSIGNED NOT NULL,
  `treasure_sequence` bigint(0) UNSIGNED NOT NULL,
  `stage_experience` bigint(0) NOT NULL,
  `stage_reward` bigint(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

 --  user_power_task 用户帮助助力任务表  --  
CREATE TABLE `user_power_task`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `publishes_user_id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `helper_user_id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_power_task_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_user_power_task_publishes_user_id`(`publishes_user_id`) USING BTREE,
  INDEX `idx_user_power_task_helper_user_id`(`helper_user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;


 --  user_publishes_helper_task 用户发布助力任务表 --  
CREATE TABLE `user_publishes_helper_task`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `user_id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `user_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `avatar` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `article` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `link` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `label` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_publishes_helper_task_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_user_publishes_helper_task_user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

